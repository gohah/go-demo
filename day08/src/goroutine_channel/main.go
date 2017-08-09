package main

import (
	"fmt"
	"time"
)

func write(c chan int) {
	for i:=0; i<100; i++ {
		c <- i
	}
}

func read(c chan int) {
	for i:=0; i<100; i++ {
		a := <- c

		fmt.Println(a);
	}
}

func main() {
	var c chan int
	c = make(chan int,10)

	go write(c)

	go read(c)

	time.Sleep(10* time.Second)
}
