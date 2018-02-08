package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	t, err := time.Parse("2006-01-02 15:04:05", "2018-01-01 00:00:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Unix())
}
