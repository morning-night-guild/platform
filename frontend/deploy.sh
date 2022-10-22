#!/bin/bash

export PROTOC_VERSION=v3.20.2

yum update

yum install wget curl

wget https://github.com/protocolbuffers/protobuf/releases/download/${PROTOC_VERSION}/protoc-$(echo ${PROTOC_VERSION} | sed 's/v//')-linux-x86_64.zip -O protobuf.zip && \
    unzip protobuf.zip -d /usr/local/bin/protobuf && \
    rm protobuf.zip && \
    chmod -R 755 /usr/local/bin/protobuf/*

export PATH=$PATH:/usr/local/bin/protobuf/

curl -sL https://rpm.nodesource.com/setup_8.x | bash -

yum install nodejs

npm install -g yarn

(cd ../api && npm install && buf generage --template buf.frontend.gen.yaml)

yarn build
