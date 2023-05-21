FROM golang:latest

# Add a specific user
RUN useradd -m Mathieu

# Install the required tools
RUN apt-get update && \
    apt-get install -y hugo && \
    apt-get install -y npm nodejs && \
    npm install -g npm@9.6.7 && \
    npm install -g markdownlint-cli@0.29.0 && \
    npm install --package-lock-only && \
    npm audit fix --package-lock-only && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Set the working directory in the container
WORKDIR /usr/src/myapp

# Switch to the new user
USER Mathieu
