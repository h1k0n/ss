#https://linuxhint.com/install_tftp_server_ubuntu/
#设置网络接口网段　将itb文件放进~/tftp
sudo ip addr add  192.168.157.2/24 dev  enx0826ae39b315
sudo ip link set enx0826ae39b315 up
sudo ip addr add  192.168.157.2/24 dev  enx0826ae39b315

#typeC作为电源，uart控制，接入电源瞬间通过按任意按键停留在uboot状态
setenv ipaddr 192.168.157.3	# 设置开发板ip
setenv serverip 192.168.157.2	# 设置服务器ip
tftpboot 0x82000000 192.168.157.3	# 内存只有80M
bootm 0x82000000
