FROM golang:latest

# Add a specific user
RUN useradd -m Mathieu

# Install the required tools
RUN apt-get update && \
    apt-get install -y hugo && \
    apt-get install -y npm nodejs && \
    npm install -g markdownlint-cli && \
    export GO111MODULE=on && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.33.0

# Set the working directory in the container
WORKDIR /usr/src/myapp

# Switch to the new user
USER Mathieu
