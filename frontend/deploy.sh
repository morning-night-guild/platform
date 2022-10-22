#!/bin/bash

yum update

yum install wget curl

export BIN="/usr/local/bin"

VERSION="v3.20.2" && \
    wget https://github.com/protocolbuffers/protobuf/releases/download/${VERSION}/protoc-$(echo ${VERSION} | sed 's/v//')-linux-x86_64.zip -O protobuf.zip && \
    unzip protobuf.zip -d /usr/local/bin/protobuf && \
    chmod -R +x "${BIN}/protobuf/*"

export PATH=$PATH:/usr/local/bin/protobuf/

VERSION="1.9.0" && \
  curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
    -o "${BIN}/buf" && \
  chmod +x "${BIN}/buf"

(cd ../api && npm install && buf generate --template buf.frontend.gen.yaml)

yarn build
