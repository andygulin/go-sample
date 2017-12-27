:: Clear Set
SET CGO_ENABLED=
SET GOOS=
SET GOARCH=

:: Build Windows
go build args.go
go build encrypt.go
go build http.go
go build json.go
go build test.go
go build time.go

:: Build Linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build args.go
go build encrypt.go
go build http.go
go build json.go
go build test.go
go build time.go