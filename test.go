package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
)

func init() {
	fmt.Println("Init...")
}

func main() {
	s := "a"
	i := 1
	b := true
	fmt.Println(s, i, b)
	fmt.Println("Hello World~")

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	fmt.Println(string(body))

	fmt.Println(os.Getpid())
}
