#!/bin/bash -e -x
apt update
apt install -yqq git curl
wget -c https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz
tar -xzf go1.11.2.linux-amd64.tar.gz
mv go go1.11.2
mv go1.11.2 /usr/local
export GOPATH=~/go
export GOROOT=/usr/local/go1.11.2
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
go get github.com/google/netstack/tcpip/sample/tun_tcp_echo
ip tuntap add user $USER mode tun tap0
ip link set tap0 up
ip addr add 172.19.0.10/12 dev tap0
#tun_tcp_echo tap0 172.19.0.4 3400
#curl 172.19.0.4:3400 -v
