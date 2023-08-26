git clone https://github.com/u-root/u-root.git
cd u-root
go build
export GOARCH=riscv64
_FLAGS=-build=bb
./u-root $_FLAGS -o abc.cpio boot core
xz --check=crc32 -9 --lzma2=dict=1MiB --stdout abc.cpio  | dd conv=sync bs=512    of=abc.initramfs.linux_riscv64.cpio.xz
