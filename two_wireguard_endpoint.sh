sudo apt-get install software-properties-common
sudo add-apt-repository ppa:wireguard/wireguard
sudo apt-get update
sudo apt-get install wireguard

wget https://launchpad.net/~wireguard/+archive/ubuntu/wireguard/+files/wireguard-tools_0.0.20181018-wg1~bionic_amd64.deb
dpkg -i wireguard-tools_0.0.20181018-wg1~bionic_amd64.deb

wg genkey | tee privatekey | wg pubkey > publickey
ip link add dev wg0 type wireguard
ip link set wg0 up
ip address add dev wg0 192.168.2.1 peer 192.168.2.2
wg set wg0 listen-port 51820 private-key ./privatekey peer ggV7EyOtvhRfnDstIAYuWDWmtNMq4dMBzJOblvx0u1c= allowed-ips 192.168.2.2 endpoint 23.252.111.86:51820


wg genkey | tee privatekey | wg pubkey > publickey
ip link add dev wg0 type wireguard
ip link set wg0 up
ip address add dev wg0 192.168.2.2 peer 192.168.2.1
wg set wg0 listen-port 51820 private-key ./privatekey peer aCssG2pmHXaewDw8nJU42Wf55yCTlpwsLQrEdH4bTl8= allowed-ips 192.168.2.1 endpoint 69.171.68.67:51820


aA6XkTvs5pDtNaXmIcadEkrb7eSZ1k6tz9VLNN+gznI=
ggV7EyOtvhRfnDstIAYuWDWmtNMq4dMBzJOblvx0u1c=


modinfo wireguard

ip link set wg0 down
ip link del dev wg0
rmmod	 wireguard

tc qdisc replace dev wg0 root noqueue


setcap cap_net_admin+eip /usr/bin/wireguard-go
