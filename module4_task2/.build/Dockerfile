FROM golang:1.15.8-alpine

# Install GCC and other required tools
RUN apk update && \
    apk add --no-cache hugo gcc musl-dev && \
    export GO111MODULE=on && \
    go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.33.0

# Set the working directory in the container
WORKDIR /usr/src/myapp
