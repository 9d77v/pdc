FROM golang:1.16.0-alpine3.13

WORKDIR /app
ENV CGO_ENABLED 0
ENV GOPROXY=https://goproxy.cn

RUN echo "http://mirrors.aliyun.com/alpine/v3.11/main/" > /etc/apk/repositories \
    && echo "http://mirrors.aliyun.com/alpine/v3.11/community/" >> /etc/apk/repositories \
    && apk add --no-cache upx

COPY go.mod /app
COPY go.sum /app

RUN go mod download \
    && rm -rf /app

