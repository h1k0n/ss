netns0="wg-test-$$-0"
netns1="wg-test-$$-1"
netns2="wg-test-$$-2"

ip netns del $netns0 2>/dev/null || true
ip netns del $netns1 2>/dev/null || true
ip netns del $netns2 2>/dev/null || true

ip netns add $netns0
ip netns add $netns1
ip netns add $netns2

ip -n $netns0 link set up dev lo

exec 3>&1
pretty() { echo -e "\x1b[32m\x1b[1m[+] ${1:+NS$1: }${2}\x1b[0m" >&3; }
maybe_exec() { if [[ $BASHPID -eq $$ ]]; then "$@"; else exec "$@"; fi; }
n0() { maybe_exec ip netns exec $netns0 "$@"; }
n1() { maybe_exec ip netns exec $netns1 "$@"; }
n2() { maybe_exec ip netns exec $netns2 "$@"; }
ip0() { pretty 0 "ip $*"; ip -n $netns0 "$@"; }
ip1() { pretty 1 "ip $*"; ip -n $netns1 "$@"; }
ip2() { pretty 2 "ip $*"; ip -n $netns2 "$@"; }

program=/usr/bin/wireguard-go
export WG_I_PREFER_BUGGY_USERSPACE_TO_POLISHED_KMOD=1
n0 $program  wg1
ip0 link set wg1 netns $netns1
n0 $program wg2
ip0 link set wg2 netns $netns2

ip1 addr add 192.168.241.1/24 dev wg1
ip1 addr add fd00::1/24 dev wg1

ip2 addr add 192.168.241.2/24 dev wg2
ip2 addr add fd00::2/24 dev wg2

pp() { pretty "" "$*"; "$@"; }
key1="$(pp wg genkey)"
key2="$(pp wg genkey)"
pub1="$(pp wg pubkey <<<"$key1")"
pub2="$(pp wg pubkey <<<"$key2")"
psk="$(pp wg genpsk)"
[[ -n $key1 && -n $key2 && -n $psk ]]

n0 wg set wg1 \
    private-key <(echo "$key1") \
    listen-port 10000 \
    peer "$pub2" \
    preshared-key <(echo "$psk") \
    allowed-ips 192.168.241.2/32,fd00::2/128
n0 wg set wg2 \
    private-key <(echo "$key2") \
    listen-port 20000 \
    peer "$pub1" \
    preshared-key <(echo "$psk") \
    allowed-ips 192.168.241.1/32,fd00::1/128

n0 wg showconf wg1
n0 wg showconf wg2

ip1 link set up dev wg1
ip2 link set up dev wg2

#ip netns exec $netns2 ping -c 10 -f -W 1 192.168.241.1

n0 wg set wg1 peer "$pub2" endpoint 127.0.0.1:20000
n0 wg set wg2 peer "$pub1" endpoint 127.0.0.1:10000

ip netns exec $netns2 ping -c 10 -f -W 1 192.168.241.1

to_kill="$(ip netns pids $netns0) $(ip netns pids $netns1) $(ip netns pids $netns2)"
[[ -n $to_kill ]] && kill $to_kill


pp ip netns del $netns1
pp ip netns del $netns2
pp ip netns del $netns0