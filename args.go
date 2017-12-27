package main

import (
	"flag"
	"fmt"
)

func main() {
	// -u=aaa -p=bbb
	u := flag.String("u", "", "username")
	p := flag.String("p", "", "password")

	flag.Parse()
	fmt.Println("u : ", *u)
	fmt.Println("p : ", *p)
}
