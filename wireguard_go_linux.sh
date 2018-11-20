#!/bin/sh
export LOG_LEVEL=debug
export WG_I_PREFER_BUGGY_USERSPACE_TO_POLISHED_KMOD=1
#wireguard-go -f wg0 #猜测加了-f不行,看代码果然是
wireguard-go wg0


wg genkey | tee privatekey | wg pubkey > publickey
wg genpsk > psk
ip addr add 192.168.241.1/24 dev wg0
ip link set up dev wg0
wg set wg0 listen-port 51823 private-key ./privatekey peer PNesgVm77WdonRj/YbVhFXmxpboCqX+Q8TYTVUyaJio= preshared-key ./psk allowed-ips 192.168.241.2/32

sysctl -w net.ipv4.ip_forward=1
iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
