#!/bin/sh
find ../ -name "*.go" | xargs gofmt -l -d -w

# LINUX
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o application  main.go

# upx压缩
upx application
# MAC
# CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o application  main.go