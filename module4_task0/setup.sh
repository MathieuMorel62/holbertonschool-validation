#!/bin/bash

# Install required packages
apt-get update && apt-get install -y git curl make wget build-essential

# Install Hugo
wget https://github.com/gohugoio/hugo/releases/download/v0.109.0/hugo_extended_0.109.0_Linux-64bit.deb
dpkg -xzf hugo_extended_0.109.0_Linux-amd64.deb
