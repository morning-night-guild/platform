FROM gcr.io/cloud-builders/go:alpine-1.18

ARG AQUA_VERSION
ARG PROTOC_VERSION

ENV GOBIN /root/go/bin

RUN apk add curl
RUN curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v1.1.0/aqua-installer | sh -s -- -i /usr/local/bin/aqua -v ${AQUA_VERSION}
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/${PROTOC_VERSION}/protoc-$(echo ${PROTOC_VERSION} | sed 's/v//')-linux-x86_64.zip -O protobuf.zip && \
    unzip protobuf.zip -d /usr/local/bin/protobuf && \
    rm protobuf.zip && \
    chmod -R 755 /usr/local/bin/protobuf/*

ENV PATH $PATH:$GOBIN:/usr/local/bin/protobuf/bin:/builder/home/.local/share/aquaproj-aqua/bin
