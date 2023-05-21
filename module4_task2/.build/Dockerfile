FROM golang:1.15.8-alpine

# Install GCC and other required tools
RUN apk update && \
    apk add --no-cache hugo gcc musl-dev && \
    export GO111MODULE=on && \
    go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.33.0

# Install npm and markdownlint
RUN apk add --no-cache npm=7.24.0 && \
    npm install -g markdownlint-cli@0.29.0

# Set the PATH to include npm
ENV PATH="/usr/local/bin:$PATH"

# Set the working directory in the container
WORKDIR /usr/src/myapp
