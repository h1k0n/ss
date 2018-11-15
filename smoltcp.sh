#!/bin/bash -e -x
apt update
apt install -yqq curl git gcc
curl https://sh.rustup.rs -sSf | sh
source $HOME/.cargo/env
git clone https://github.com/m-labs/smoltcp.git
cd smoltcp/
ip tuntap add name tap0 mode tap user $USER
ip link set tap0 up
ip addr add 192.168.69.100/24 dev tap0
cargo build --example httpclient
#./target/debug/examples/httpclient tap0 93.184.216.34 http://example.org/
#tun_tcp_echo -tap tap0 192.168.69.3 3400
#tcpdump -i tap0

#launch tcp endpoint
#tun_tcp_echo -tap tap0 192.168.69.3 80

#send sync
#go get github.com/grahamking/latency
#latency 192.168.69.3

#observe packets
#tcpdump -i tap0

#observe packets
#iptables -t raw -A OUTPUT -o tap0 -p tcp -j LOG --log-prefix "IPTABLES raw-OUT: "
#iptables -t mangle -A OUTPUT -o tap0 -p tcp -j LOG --log-prefix "IPTABLES mangle-OUT: "
#iptables -t nat -A OUTPUT -o tap0 -p tcp -j LOG --log-prefix "IPTABLES nat-OUT: "
#iptables -t filter -A OUTPUT -o tap0 -p tcp -j LOG --log-prefix "IPTABLES filter-OUT: "
#tail -f /var/log/syslog

#i will try to use gopacket to send packet next time.
#try to use smoltcp netstack example server
#try to figure out nat output chain
#why first packet's checksum incorrec
#cargo run --example server -- tap0
#latency  -p 6969 192.168.69.1

#if i don't indicate interface,latancy will use default interface, it lead to checksum error.
#cargo run --example server -- tap0
#latency -i tap0 -p 6969 192.168.69.1

#tun_tcp_echo -tap tap0 192.168.69.3 80
#latency -i tap0 192.168.69.3

#when use latency and tun_tcp_echo with no -i tap0, why second packet's checksum correct automaticaly?
