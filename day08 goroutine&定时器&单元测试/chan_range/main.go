package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 1000)

	for i := 0; i < 1000; i++ {
		ch <- i
	}

	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
