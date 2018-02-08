package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Hostname())
	fmt.Println(os.Getpid())
	fmt.Println(os.TempDir())
	envs := os.Environ()
	for _, env := range envs {
		fmt.Println(env)
	}
}
