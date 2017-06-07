package main

import(
	"fmt"
)

var pipe chan int

func main() {

	//pipe = make(chan int, 1)
	//
	//go test_signal(1,2)
	//
	//c :=<- pipe
	//
	//fmt.Println(c)

	a,b := calc(1,2)

	fmt.Println(a,b)
}
