FROM gcr.io/cloud-builders/gcloud

ARG KO_VERSION
ARG GO_VERSION

RUN apt update && apt install -y wget && \
    wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
ENV GOBIN /usr/local/go/bin
ENV PATH $PATH:$GOBIN
RUN go install github.com/google/ko@${KO_VERSION}
