#!/bin/bash

# Install required packages
sudo apt-get update && sudo apt-get install -y curl make git zip wget npm

# Install Hugo
sudo curl -L https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.deb -o hugo.deb
sudo apt install ./hugo.deb
export PATH=/usr/local/bin:$PATH

#remove file after installation
sudo rm hugo.deb

# Uninstall go
sudo apt-get remove golang-go
sudo rm -rf /usr/local/go

# remove dist directory
sudo rm -rf dist/ 2> /dev/null
