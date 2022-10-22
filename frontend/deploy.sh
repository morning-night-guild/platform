#!/bin/bash

export PROTOC_VERSION=v3.20.2

yum update

yum install wget curl

wget https://github.com/protocolbuffers/protobuf/releases/download/${PROTOC_VERSION}/protoc-$(echo ${PROTOC_VERSION} | sed 's/v//')-linux-x86_64.zip -O protobuf.zip && \
    unzip protobuf.zip -d /usr/local/bin/protobuf && \
    rm protobuf.zip && \
    chmod -R 755 /usr/local/bin/protobuf/*

export PATH=$PATH:/usr/local/bin/protobuf/

BIN="/usr/local/bin" && \
VERSION="1.9.0" && \
  curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
    -o "${BIN}/buf" && \
  chmod +x "${BIN}/buf"

(cd ../api && npm install && buf generage --template buf.frontend.gen.yaml)

yarn build
