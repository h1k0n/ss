#out of memory
sudo dd if=/dev/zero of=/tmp/.swapfile bs=1024 count=2097152
sudo chmod 600 /tmp/.swapfile
sudo mkswap /tmp/.swapfile
sudo swapon /tmp/.swapfile

# configuaration.nix
services.openssh = {  
   enable = true;  
   ports = [XX];
};
vim
wget

#upgrade
sudo nixos-rebuild switch --upgrade

nix-channel --add https://nixos.org/channels/nixos-23.11 nixos
#Channels are set per user. This means that running nix-channel --add as a non root user (or without sudo) will not affect configuration in /etc/nixos/configuration.nix
nixos-rebuild switch --upgrade

sudo nix-collect-garbage -d
