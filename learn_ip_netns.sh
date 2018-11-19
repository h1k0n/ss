#!/bin/bash
apt update
apt install bridge-utils
ip netns add ns1

ip netns exec ns1 ifconfig -a


ip link add name ns1-nic type veth peer name ns1-vnic
ip link set ns1-vnic netns ns1
ip netns exec ns1 ip link set ns1-vnic name eth0
ip netns exec ns1 ip addr add 10.0.0.100/24 dev eth0
ip netns exec ns1 ip link set eth0 up

brctl addbr testbr
brctl addif testbr ns1-nic

ip link set ns1-nic up
ip addr add 10.0.0.1/24 dev testbr

ip link set testbr up
ping 10.0.0.100
ip netns exec ns1 ping 10.0.0.1