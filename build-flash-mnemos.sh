# build github.com/jspngh/oreboot 
# nix-shell nix-shell
cargo make --release --variant lichee --supervisor #gen oreboot-nezha-bt0.bin
sudo dd if=oreboot-nezha-bt0.bin of=/dev/sdb bs=1024 seek=8 conv=sync

# nix-shell latest-rust.nix
just build-d1 lichee-rv
just build-d1 mq-pro
sudo dd if=mnemos-lichee-rv.bin of=/dev/sdb bs=1024 seek=40 conv=sync
sudo dd if=mnemos-mq-pro.bin of=/dev/sdb bs=1024 seek=40 conv=sync
