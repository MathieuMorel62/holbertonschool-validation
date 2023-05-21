FROM golang:1.15.8-alpine

# Install the required tools
RUN apk update && \
    apk add --no-cache hugo && \
    export GO111MODULE=on && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.33.0

# Set the working directory in the container
WORKDIR /usr/src/myapp
