# sed modiy arch/riscv/configs/d1_defconfig CONFIG_INITRAMFS_SOURCE="/tmp/initramfs.linux_riscv64.cpio.xz"
#export CCPREFIX=/home/ha1kang/riscv64-glibc/riscv/bin/riscv64-unknown-linux-gnu-   failed to boot: supervisor illegal instruction
export CCPREFIX=/usr/bin/riscv64-linux-gnu-
make ARCH=riscv CROSS_COMPILE=$CCPREFIX clean
make ARCH=riscv CROSS_COMPILE=$CCPREFIX d1_defconfig
make ARCH=riscv CROSS_COMPILE=$CCPREFIX 

result:
  DTC     arch/riscv/boot/dts/allwinner/sun20i-d1-nezha.dtb
  DTC     arch/riscv/boot/dts/allwinner/sun20i-d1-lichee-rv.dtb
  DTC     arch/riscv/boot/dts/allwinner/sun20i-d1-lichee-rv-dock.dtb
  DTC     arch/riscv/boot/dts/allwinner/sun20i-d1-lichee-rv-jeadock.dtb
  DTC     arch/riscv/boot/dts/allwinner/sun20i-d1-dongshan-nezha-stu.dtb
  DTC     arch/riscv/boot/dts/allwinner/sun20i-d1-mangopi-rmii.dtb
  
  OBJCOPY arch/riscv/boot/Image
  GZIP    arch/riscv/boot/Image.gz
  Kernel: arch/riscv/boot/Image.gz is ready
  
