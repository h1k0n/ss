/*
based on
https://discourse.nixos.org/t/how-can-i-set-up-my-rust-programming-environment/4501/9
*/
let
  rust_overlay = import (builtins.fetchTarball "https://github.com/oxalica/rust-overlay/archive/master.tar.gz");
  pkgs = import <nixpkgs> { overlays = [ rust_overlay ]; };
  rustVersion = "latest";
  #rustVersion = "1.62.0";
  rust = pkgs.rust-bin.nightly."2023-03-16".default.override {
    targets = [ "riscv64imac-unknown-none-elf" ];
    extensions = [
      "rust-src" # for rust-analyzer
      "llvm-tools-preview"
    ];
  };
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
  ]);
  shellHook = ''
    export PATH="$PATH:$HOME/.cargo/bin"
  '';
  RUST_BACKTRACE = 1;
}
