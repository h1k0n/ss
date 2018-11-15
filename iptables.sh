#iptables -t raw -A OUTPUT -o tap0 -p tcp -j LOG --log-prefix "IPTABLES raw-OUT: "
#iptables -t mangle -A OUTPUT -o tap0 -p tcp -j LOG --log-prefix "IPTABLES mangle-OUT: "
#iptables -t filter -A OUTPUT -o tap0 -p tcp -j LOG --log-prefix "IPTABLES filter-OUT: "


iptables -t nat -N forward_to_mypc
iptables -t nat -A forward_to_mypc -p tcp -j LOG --log-prefix "IPTABLES nat-OUT: "
iptables -t nat -A forward_to_mypc -p tcp -j DNAT --to 192.168.69.3:80

iptables -t nat -I OUTPUT -p tcp --dport 3400 -j forward_to_mypc

tun_tcp_echo -tap tap0 192.168.69.3 80

latency -i tap0 -p 3400 192.168.69.1




iptables -t nat -D OUTPUT 1
iptables -t nat --flush forward_to_mypc
iptables -t nat -X forward_to_mypc


iptables -t nat -N forward_to_mypc
iptables -t nat -A forward_to_mypc -p tcp -j LOG --log-prefix "IPTABLES nat-OUT: "
iptables -t nat -A forward_to_mypc -p tcp -j DNAT --to 192.168.69.1:6969

iptables -t nat -I OUTPUT -p tcp --dport 3400 -j forward_to_mypc

cargo run --example server -- tap0

latency -i tap0 -p 3400 192.168.69.1

