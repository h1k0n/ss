# 扩大加载jspngh/oreboot从tf卡的加载长度，重新编译，获取oreboot-nezha-bt0.bin
#　编译michaelengel/xv6-d1，获取kernel.bin
sudo dd if=oreboot-nezha-bt0.bin of=/dev/sdb bs=1024 seek=8 conv=sync
sudo dd if=kernel.bin of=/dev/sdb bs=1024 seek=40 conv=sync
