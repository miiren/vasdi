#!/bin/sh
find ./ -name "*.go" | xargs gofmt -l -d -w

# MAC
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dev-application main.go

# win
# CGO_ENABLED=0 GOOS=windows  GOARCH=amd64 go build -o dev-application main.go

# upx压缩
upx dev-application