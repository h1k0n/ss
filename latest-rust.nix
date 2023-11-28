/*
based on
https://discourse.nixos.org/t/how-can-i-set-up-my-rust-programming-environment/4501/9
*/
let
  rust_overlay = import (builtins.fetchTarball "https://github.com/oxalica/rust-overlay/archive/master.tar.gz");
  pkgs = import <nixpkgs> { overlays = [ rust_overlay ]; };
  rustVersion = "latest";
  #rustVersion = "1.62.0";
  rust = pkgs.rust-bin.selectLatestNightlyWith (toolchain: toolchain.default.override {
    targets = [ "riscv64imac-unknown-none-elf" ];
    # arceos使用 targets = [ "riscv64gc-unknown-none-elf" ];
    extensions = [
      "rust-src" # for rust-analyzer
      "llvm-tools-preview"
    ];
  });
in
pkgs.mkShell {
  buildInputs = [
    rust
  ] ++ (with pkgs; [
    rust-analyzer
    pkg-config
    dtc
    qemu
    # other dependencies
    #gtk3
    #wrapGAppsHook
    just
  ]);
  shellHook = ''
    export PATH="$PATH:$HOME/.cargo/bin"
  '';
  RUST_BACKTRACE = 1;
}
