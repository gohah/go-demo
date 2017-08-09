package main

import (
	"time"
	"fmt"
)
//每隔一秒执行一次
func main() {
	t := time.NewTicker(time.Second)

	for v := range t.C {
		fmt.Println(v)
	}
}
