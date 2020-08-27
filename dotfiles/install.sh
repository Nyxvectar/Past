#!/bin/bash

# @Encode : UTF-8
# @Author : Nyxvectar
# @Repo   : dotfiles
# @File   : install
# @Time   : 7/11/24
# @IDE    : GoLand

sudo mkdir -p /usr/share/fonts/
sudo mkdir -p /etc/fonts/
sudo mkdir -p /root/.config/
mkdir -p ~/.config/
mkdir -p ~/.config/Code/User/

sudo cp annex/fonts/'Proxima Nova.ttf' /usr/share/fonts/
sudo cp annex/fonts/'Proxima Nova.ttf' /etc/fonts/
sudo cp -r fish /root/.config/

cp -r alacritty ~/.config/
cp -r fish ~/.config/
cp -r fontconfig ~/.config/
cp -r hypr ~/.config/
cp -r kitty ~/.config/
cp -r mako ~/.config/
cp -r rofi ~/.config/
cp -r waybar ~/.config/
cp -r fcitx/themes ~/.local/share/fcitx5/
cp -r annex/vscode/settings.json ~/.config/Code/User/

echo "Done."