#!/bin/bash

# Update packages and install hugo and make
apt-get update && apt-get install -y hugo make

# Run make build
make build
