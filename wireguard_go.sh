
export LOG_LEVEL=debug
export WG_I_PREFER_BUGGY_USERSPACE_TO_POLISHED_KMOD=1
./wireguard-go -f wg0

./wg genkey | tee privatekey | ./wg pubkey > publickey

ip link set wg0 up
ip address add dev wg0 192.168.2.1 peer 192.168.2.2
./wg set wg0 listen-port 51820 private-key ./privatekey peer 2vh1CwMycQIAdCxMI/u2JuAvvETVrq4dO2GgKovySEg= allowed-ips 192.168.2.2 endpoint 23.252.111.86:51820


ip link set wg0 up
ip address add dev wg0 192.168.2.2 peer 192.168.2.1
./wg set wg0 listen-port 51820 private-key ./privatekey peer Ivoaf0v2k62ieTqAjxFStLpzX1tIxFnKqEos+18bHh0= allowed-ips 192.168.2.1 endpoint 120.77.69.87:51820


iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE


sudo apt-get install libmnl-dev libelf-dev linux-headers-$(uname -r) build-essential pkg-config
git clone https://git.zx2c4.com/WireGuard
cd WireGuard/src
make
mkdir ~/wg && cp tools/wg $_ 
cd ~/wg


wget -c https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz
tar -xzf go1.11.2.linux-amd64.tar.gz
mv go go1.11.2
mv go1.11.2 /usr/local
export GOROOT=/usr/local/go1.11.2
export PATH=$GOROOT/bin:$PATH
printf 'package main\nconst UseTheKernelModuleInstead = 0xdeadbabe\n' >ireallywantobuildon_linux.go
go build
mkdir ~/wireguard && cp wireguard-go $_ 

