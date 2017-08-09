package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()

	var m map[string]int
	m["stu"] = 100
}

func calc() {
	for {
		fmt.Println("i'm calc")
		time.Sleep(time.Second)
	}
}

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	go test()
	for i := 0; i < 2; i++ {
		go calc()
	}

	time.Sleep(time.Second * 10000)
}
