package main

import (
	"time"
	"fmt"
)
//一次定时
func main() {
	select {
	case <- time.After(time.Second):
		fmt.Println("after one second")
	}
}
