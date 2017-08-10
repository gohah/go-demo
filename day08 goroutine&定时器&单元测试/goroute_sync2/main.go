package main

import "fmt"

func send(ch chan int, exitChan chan struct{}) {

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	var a struct{}
	exitChan <- a
}

func recv(ch chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}

	var a struct{}
	exitChan <- a
}

func main() {
	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)

	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
}
