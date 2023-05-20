#!/bin/bash

# Install required packages
apt-get update && apt-get install -y curl make npm wget 

# Install Hugo
curl -LO https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.tar.gz
tar -xzf hugo_extended_0.84.0_Linux-64bit.tar.gz
