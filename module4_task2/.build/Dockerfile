FROM golang:1.15.8-alpine

# Install GCC and other required tools
# hadolint ignore=DL3018
RUN apk update && \
    apk add --no-cache --update hugo gcc musl-dev && \
    export GO111MODULE=on && \
    go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.33.0 && \
    rm -rf /var/cache/apk/*

# Install Node.js and npm
# hadolint ignore=DL3018,DL3019
RUN apk add --no-cache --update nodejs npm && \
    rm -rf /var/cache/apk/*

# Install markdownlint
# hadolint ignore=DL3016
RUN npm install -g markdownlint-cli

# Install Hadolint
RUN wget --progress=dot:giga https://github.com/hadolint/hadolint/releases/download/v2.6.0/hadolint-Linux-x86_64 && \
    chmod +x hadolint-Linux-x86_64 && \
    mv hadolint-Linux-x86_64 /usr/local/bin/hadolint

# Install Container Structure Test
RUN wget --progress=dot:giga -O /usr/local/bin/container-structure-test \
    https://storage.googleapis.com/container-structure-test/v1.10.0/container-structure-test-linux-amd64 && \
    chmod +x /usr/local/bin/container-structure-test

# Set the working directory in the container
WORKDIR /usr/src/myapp
