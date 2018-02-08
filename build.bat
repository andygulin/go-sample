:: Build Windows
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o bin\args.exe args.go
go build -o bin\encrypt.exe encrypt.go
go build -o bin\goroutine.exe goroutine.go
go build -o bin\http.exe http.go
go build -o bin\http2.exe http2.go
go build -o bin\json.exe json.go
go build -o bin\mongo.exe mongo.go
go build -o bin\mysql.exe mysql.go
go build -o bin\os.exe os.go
go build -o bin\redis.exe redis.go
go build -o bin\reflection.exe reflection.go
go build -o bin\test.exe test.go
go build -o bin\time.exe time.go

:: Build Linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o bin\args args.go
go build -o bin\encrypt encrypt.go
go build -o bin\goroutine goroutine.go
go build -o bin\http http.go
go build -o bin\http2 http2.go
go build -o bin\json json.go
go build -o bin\mongo mongo.go
go build -o bin\mysql mysql.go
go build -o bin\os os.go
go build -o bin\redis redis.go
go build -o bin\reflection reflection.go
go build -o bin\test test.go
go build -o bin\time time.go