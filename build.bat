:: Clear Set
SET CGO_ENABLED=
SET GOOS=
SET GOARCH=

:: Build Windows
go build test.go

:: Build Linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build test.go