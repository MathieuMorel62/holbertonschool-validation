#!/bin/bash

# Install required packages
apt-get update && apt-get install -y curl make

# Install Hugo
curl -LO https://github.com/gohugoio/hugo/releases/download/v0.111.3/hugo_0.111.3_Linux-ARM64.tar.gz
tar -xzf hugo_0.111.3_Linux-ARM64.tar.gz -C /usr/local/bin/
rm hugo_0.111.3_Linux-ARM64.tar.gz
export PATH=/usr/local/bin:$PATH

# Run make build
make build
