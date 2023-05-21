FROM golang:1.15.8-alpine

# Install GCC and other required tools
RUN apk update && \
    apk add --no-cache hugo gcc musl-dev && \
    export GO111MODULE=on && \
    go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.33.0

# Install npm and markdownlint
RUN apk add --no-cache npm && \
    npm install -g markdownlint-cli

# Install container-structure-test
ENV CONTAINER_STRUCTURE_TEST_VERSION=1.10.0
RUN wget -O /usr/local/bin/container-structure-test https://storage.googleapis.com/container-structure-test/v${CONTAINER_STRUCTURE_TEST_VERSION}/container-structure-test-linux-amd64 && \
    chmod +x /usr/local/bin/container-structure-test

# Set the PATH to include npm
ENV PATH="/usr/local/bin:$PATH"

# Set the working directory in the container
WORKDIR /usr/src/myapp
