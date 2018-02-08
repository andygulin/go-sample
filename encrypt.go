package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte("123456"))
	md5Str := hex.EncodeToString(md5Ctx.Sum(nil))
	fmt.Println(md5Str)

	base64Str := base64.StdEncoding.EncodeToString([]byte("123456"))
	fmt.Println(base64Str)

	bytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}
