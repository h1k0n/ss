#!/bin/bash
apt update
apt install -yqq gcc make git nmap iperf3


wget -c https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz
tar -xzf go1.11.2.linux-amd64.tar.gz
mv go go1.11.2
mv go1.11.2 /usr/local
export GOPATH=~/go
export GOROOT=/usr/local/go1.11.2
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

wget "https://github.com/WireGuard/wireguard-go/archive/0.0.20181018.tar.gz"
tar -xvf 0.0.20181018.tar.gz
cd wireguard-go-0.0.20181018
sed -i "s{ifeq{ifneq{" Makefile
make
make install


wget https://launchpad.net/~wireguard/+archive/ubuntu/wireguard/+files/wireguard-tools_0.0.20181018-wg1~bionic_amd64.deb
dpkg -i wireguard-tools_0.0.20181018-wg1~bionic_amd64.deb

export WG_I_PREFER_BUGGY_USERSPACE_TO_POLISHED_KMOD=1
cd tests/
./netns.sh /usr/bin/wireguard-go

