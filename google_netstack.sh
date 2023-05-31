#!/bin/bash -e -x
apt update
apt install -yqq git curl
wget -c https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
tar -xzf go1.11.5.linux-amd64.tar.gz
mv go go1.11.5
mv go1.11.5 /usr/local
export GOPATH=~/go
export GOROOT=/usr/local/go1.11.5
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
go get github.com/google/netstack/tcpip/sample/tun_tcp_echo
ip tuntap add user $USER mode tun tun5
ip link set tun5 up
ip addr add 192.168.57.10/12 dev tun5  //指定路由
#tun_tcp_echo tap0 192.168.57.11 7000 //指定用户协议栈的echo服务地址
#curl 192.168.57.11:7000 -v   //请求echo服务

#新的tun_tcp_echo: https://github.com/google/gvisor/tree/master/pkg/tcpip/sample/tun_tcp_echo
