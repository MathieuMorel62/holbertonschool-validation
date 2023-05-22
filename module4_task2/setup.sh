#!/bin/bash

# Update package lists and installation of required packages
apt-get update && apt-get install -y hugo make npm curl wget
