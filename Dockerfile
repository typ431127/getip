ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir -p /api
WORKDIR /api
ENV GOPROXY=https://goproxy.cn,direct

COPY . .
RUN go build -o ./app main.go

FROM alpine:3.14

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/app .

EXPOSE 8080

ENTRYPOINT ["./app"]