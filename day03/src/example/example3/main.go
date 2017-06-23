package main

import (
	"time"
	"fmt"
)

func test() {
	time.Sleep(time.Millisecond*100)
}

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()

	fmt.Println((end-start)/1000)
}
