FROM golang:1.15 AS builder
MAINTAINER Andrey Martynenko <ya.lergor@yandex.ru>

# enable Go modules support
ENV GO111MODULE=on
ENV GOOS=linux
ENV CGO_ENABLED=1

WORKDIR /app

COPY . .

# install go mod, compiler protobuf & spinner
RUN go mod download \
    && go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o hello-page-and-return-name cmd/hello-page-and-return-name/main.go

FROM alpine:3.4
WORKDIR /app

COPY --from=builder /app/hello-page-and-return-name /usr/local/bin

CMD ["hello-page-and-return-name"]
