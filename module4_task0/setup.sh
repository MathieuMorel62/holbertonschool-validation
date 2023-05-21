#!/bin/bash

# Install required packages
apt-get update && apt-get install -y curl make git zip wget npm

# Install Hugo
sudo curl -LO https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.tar.gz
tar -xzf hugo_extended_0.84.0_Linux-64bit.tar.gz -C /usr/local/bin/
sudo rm hugo_extended_0.84.0_Linux-64bit.tar.gz


# Uninstall go
sudo apt-get remove golang-go
rm -rf /usr/local/go

# remove dist directory
rm -rf dist/ 2> /dev/null
