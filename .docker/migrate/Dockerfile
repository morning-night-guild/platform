ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION}-alpine

WORKDIR /app

RUN apk update && apk upgrade && \
    apk add build-base && \
    go install github.com/cosmtrek/air@latest

WORKDIR /app/cmd/migrate

CMD ["air"]
