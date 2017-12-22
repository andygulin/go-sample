#!/usr/bin/env bash
# Build Linux
go build test.go
# Build Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build test.go