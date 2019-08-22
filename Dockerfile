FROM golang:1.12.9-buster

ENV CGO_ENABLED=0

ENV GOOS=linux

ENV GOARCH=amd64

ENV GOPROXY=https://mirrors.aliyun.com/goproxy/

RUN go get github.com/derekparker/delve/cmd/dlv