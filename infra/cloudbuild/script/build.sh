#!/bin/sh

apt update && apt install -y wget curl zip gcc

# install go
wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz 

tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz

# install aqua
curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v1.1.0/aqua-installer | sh -s -- -i /usr/local/bin/aqua -v v${AQUA_VERSION}

# install protobuf
wget https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip -O protobuf.zip

unzip protobuf.zip -d /usr/local/bin/protobuf

chmod -R +x /usr/local/bin/protobuf/*

# export path
export GOBIN=/usr/local/go/bin

export PATH=$PATH:$GOBIN:usr/local/bin/protobuf/bin:/builder/home/.local/share/aquaproj-aqua/bin

# setup tools
go install github.com/google/ko@${KO_VERSION}

# build
(cd api && aqua i && buf generate --template buf.backend.gen.yaml)

(cd backend && \
    aqua i && \
    go generate ./... && \
    wire gen ./app/${SERVICE_NAME} && \
    SOURCE_DATE_EPOCH=$(date +%s) KO_DOCKER_REPO=${GCP_REPO}/${SERVICE_NAME} ko build --sbom=none --bare ./app/${SERVICE_NAME})
