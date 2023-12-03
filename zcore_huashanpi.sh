ZCOREROOT=~/zCore

git clone https://github.com/rcore-os/rcore-fs
cd rcore-fs/rcore-fs-fuse/
cargo build
export PATH=$PATH:/home/xxx/Documents/xxx/rcore-fs/target/debug #export rcore-fs-fuse path

cd $ZCOREROOT
nix flake init
nix flake update
nix develop .#default
cargo rootfs --arch riscv64
make image ARCH=riscv64
cd $ZCOREROOT/zCore
make build LINUX=1 ARCH=riscv64 PLATFORM=cv1811

