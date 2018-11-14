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
