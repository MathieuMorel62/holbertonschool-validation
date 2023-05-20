#!/bin/bash

# Install required packages
sudo apt-get update && sudo apt-get install -y git curl make wget build-essential

# Download Hugo
wget https://github.com/gohugoio/hugo/releases/download/v0.109.0/hugo_extended_0.109.0_Linux-64bit.tar.gz

# Extract the file
tar -xzf hugo_extended_0.109.0_Linux-64bit.tar.gz
