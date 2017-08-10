package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	var p *int
	var a int
	p = &a
	*p = 0

	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()
	_, err = c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
}
