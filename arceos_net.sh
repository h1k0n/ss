# 网卡驱动 https://github.com/orgs/rcore-os/discussions/30
# 硬件支持 https://github.com/orgs/rcore-os/discussions/24
# elliott10 Luoyuan Xiao 萧络元
# zxgsn 张轩铬
# yuoo655 CL  陈乐
# ZhiyuanSue ZY
# https://github.com/rcore-os/arceos/discussions/120
# 星光派2驱动适配负责人: CL陈乐；
# 华山派 + 通用PHY 的网络驱动负责人: ZY + RISCV开发群组爱好者；
# T-HEAD C910开发板适配负责人: XLY萧络元
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
