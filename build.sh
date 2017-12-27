#!/usr/bin/env bash

# Build Linux
CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
go build args.go
go build encrypt.go
go build http.go
go build json.go
go build test.go
go build time.go

# Build Windows
CGO_ENABLED=0
GOOS=windows
GOARCH=amd64
go build args.go
go build encrypt.go
go build http.go
go build json.go
go build test.go
go build time.go