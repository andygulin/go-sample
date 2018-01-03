package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	con, err := redis.Dial("tcp", "192.168.209.128:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()

	con.Do("SET", "name", "aaa")
	name, err := redis.String(con.Do("GET", "name"))
	fmt.Println(name)
}
