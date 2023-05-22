FROM golang:1.15.8-alpine

# Install GCC and other required tools
# hadolint ignore=DL3018,DL4006,SC2046
RUN apk update && \
    apk add --no-cache --update curl hugo gcc musl-dev && \
    export GO111MODULE=on && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin && \
    rm -rf /var/cache/apk/*

# Install Node.js and npm
# hadolint ignore=DL3018,DL3019
RUN apk add --no-cache --update nodejs npm && \
    rm -rf /var/cache/apk/*

# Install markdownlint
# hadolint ignore=DL3016
RUN npm install -g markdownlint-cli

# Install hadolint
RUN curl -sLo /usr/local/bin/hadolint \
    https://github.com/hadolint/hadolint/releases/download/v2.6.0/hadolint-Linux-x86_64 && \
    chmod +x /usr/local/bin/hadolint

# Install Container Structure Test
# hadolint ignore=DL4001,DL3018
RUN apk add --no-cache --update curl \
    && curl -LO https://storage.googleapis.com/container-structure-test/latest/container-structure-test-linux-amd64 \
    && chmod +x container-structure-test-linux-amd64 \
    && mv container-structure-test-linux-amd64 /usr/local/bin/container-structure-test

# Set the working directory in the container
WORKDIR /usr/src/myapp

# Add the markdownlint configuration file
COPY module4_task2/.build/cst.yml /usr/src/myapp/.build/cst.yml
