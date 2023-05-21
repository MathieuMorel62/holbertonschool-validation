FROM golang:1.15.8-alpine

# Install GCC and other required tools
RUN apk update && \
    apk add --no-cache hugo gcc musl-dev && \
    export GO111MODULE=on && \
    go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.33.0

# Install markdownlint
RUN npm install -g markdownlint-cli

# Set the PATH to include markdownlint
ENV PATH="/root/.npm-global/bin:$PATH"

# Set the working directory in the container
WORKDIR /usr/src/myapp
