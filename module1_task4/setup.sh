#!/bin/bash

# Install required packages, Hugo and Make
apt-get update && apt-get install -y curl hugo make

# Run make build
make build
