FROM golang:1.15.8-alpine

# Install GCC and other required tools
# hadolint ignore=DL3018
RUN apk --no-cache add hugo gcc musl-dev && \
    export GO111MODULE=on && \
    go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.33.0

# Install Node.js and npm
RUN apk --no-cache add nodejs=12.18.4-r0 npm=12.18.4-r0

# Install markdownlint
# hadolint ignore=DL3016
RUN npm install -g markdownlint-cli

# Set the working directory in the container
WORKDIR /usr/src/myapp

