#!/bin/bash

cd ~
if [ ! -d ~/.oh-my-zsh ]; then
    sudo cp -r /root/.oh-my-zsh $HOME/.oh-my-zsh
    sudo chown -R $USER:$USER $HOME/.oh-my-zsh
fi

if [ ! -d ~/.nvm ]; then
    sudo cp -r /root/.nvm $HOME/.nvm
    sudo chown -R $USER:$USER $HOME/.nvm
fi

if [ ! -d ~/.local/share/fonts ]; then
    mkdir -p ~/.local/share
    sudo cp -r /root/.local/share/fonts $HOME/.local/share/fonts
    sudo chown -R $USER:$USER $HOME/.local
fi

if [ ! -f ~/.zshrc ]; then
    sudo cp /root/.zshrc $HOME/.zshrc
    sudo chown -R $USER:$USER $HOME/.zshrc
fi

if [ ! -f ~/.p10k.zsh ]; then
    sudo cp /root/.p10k.zsh $HOME/.p10k.zsh
    sudo chown -R $USER:$USER $HOME/.p10k.zsh
fi