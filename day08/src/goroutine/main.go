package main

import (
	"fmt"
	//"time"
	"runtime"
)

func test() {
	for {
		fmt.Println("1")
	}
}

func main() {

	num := runtime.NumCPU()

	runtime.GOMAXPROCS(num)

	fmt.Println(num);

	//go test()
	//time.Sleep(10000000)
}
