package main

import (
	"runtime"
	"fmt"
)

func say(s string) {
	for i:=0;i<5;i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
