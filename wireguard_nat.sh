netns0="wg-test-$$-0"
netns1="wg-test-$$-1"
netns2="wg-test-$$-2"

ip netns del $netns0 2>/dev/null || true
ip netns del $netns1 2>/dev/null || true
ip netns del $netns2 2>/dev/null || true

ip netns add $netns0
ip netns add $netns1
ip netns add $netns2

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
n1 $program wg1
n2 $program wg2

pp() { pretty "" "$*"; "$@"; }
key1="$(pp wg genkey)"
key2="$(pp wg genkey)"
pub1="$(pp wg pubkey <<<"$key1")"
pub2="$(pp wg pubkey <<<"$key2")"
psk="$(pp wg genpsk)"
[[ -n $key1 && -n $key2 && -n $psk ]]

configure_peers() {

	ip1 addr add 192.168.241.1/24 dev wg1
	ip1 addr add fd00::1/24 dev wg1

	ip2 addr add 192.168.241.2/24 dev wg2
	ip2 addr add fd00::2/24 dev wg2

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
	sleep 1
}
configure_peers

ip0 link add vethrc type veth peer name vethc
ip0 link add vethrs type veth peer name veths
ip0 link set vethc netns $netns1
ip0 link set veths netns $netns2
ip0 link set vethrc up
ip0 link set vethrs up
ip0 addr add 192.168.1.1/24 dev vethrc
ip0 addr add 10.0.0.1/24 dev vethrs

ip1 addr add 192.168.1.100/24 dev vethc
ip1 link set vethc up

ip1 route add default via 192.168.1.1

ip2 addr add 10.0.0.100/24 dev veths
ip2 link set veths up

waitiface() { pretty "${1//*-}" "wait for $2 to come up"; ip netns exec "$1" bash -c "while [[ \$(< \"/sys/class/net/$2/operstate\") != up ]]; do read -t .1 -N 0 || true; done;"; }
waitiface $netns0 vethrc
waitiface $netns0 vethrs
waitiface $netns1 vethc
waitiface $netns2 veths

n0 bash -c 'printf 1 > /proc/sys/net/ipv4/ip_forward'
n0 iptables -t nat -A POSTROUTING -s 192.168.1.0/24 -d 10.0.0.0/24 -j SNAT --to 10.0.0.1
n0 wg set wg1 peer "$pub2" endpoint 10.0.0.100:20000 persistent-keepalive 1
n1 ping 192.168.241.2
n2 ping 192.168.241.1


n0 iptables -t nat -F
ip0 link del vethrc
ip0 link del vethrs
ip1 link del wg1
ip2 link del wg2















##
ip netns exec $netns0 ip addr
ip netns exec $netns1 ip addr
ip netns exec $netns2 ip addr

ip netns exec $netns1 wg showconf wg1
ip netns exec $netns2 wg showconf wg2
