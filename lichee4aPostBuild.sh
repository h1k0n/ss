#https://github.com/h1k0n/nixos-licheepi4a@8535396
nix build .#uboot -L --show-trace
nix build .#sdImage -L --show-trace

# flash u-boot into spl partition
sudo fastboot flash ram u-boot-with-spl.bin
sudo fastboot reboot
# flash uboot partition
sudo fastboot flash uboot u-boot-with-spl.bin

cp  result/sd-image/nixos-licheepi4a-sd-image-*-riscv64-linux.img  nixos-lp4a.img
chmod 644 nixos-lp4a.img
dd if=/dev/zero bs=1M count=16 >> nixos-lp4a.img
sudo losetup --find --partscan nixos-lp4a.img

# check rootfs's status, it's broken.
sudo fsck /dev/loop0p2

echo w | sudo fdisk /dev/loop0
# increase the rootfs's partition size & file system size
nix shell nixpkgs#cloud-utils
sudo growpart /dev/loop0 2

# check rootfs's status again, it should be normal now.
sudo fsck /dev/loop0p2

# umount the image file
sudo losetup -d /dev/loop0

# please replace `/dev/sdX` with your SD card's device name
sudo dd if=nixos-lp4a.img of=/dev/sdX bs=4M status=progress

# resizepart 2 to 100%
sudo parted /dev/sdX
