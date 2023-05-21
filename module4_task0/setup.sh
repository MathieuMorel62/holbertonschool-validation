#!/bin/bash

# Install required packages
sudo apt-get update && apt-get install -y curl make wget npm

# Install Hugo
sudo wget https://github.com/gohugoio/hugo/releases/download/v0.87.0/hugo_extended_0.87.0_Linux-64bit.deb
sudo dpkg -i hugo_extended_0.87.0_Linux-64bit.deb

sudo npm install -g markdownlint-cli markdown-link-check
