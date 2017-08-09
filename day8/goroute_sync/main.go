package main

import (
	"fmt"
)

func calc(taskChan chan int, resChan chan int,
	exitChan chan bool) {
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			resChan <- v
		}
	}

	fmt.Println("exit")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	go func() {
		for i := 0; i < 10000; i++ {
			intChan <- i
		}

		close(intChan)
	}()

	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan, exitChan)
	}

	//等待所有计算的goroutine全部退出
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
			fmt.Println("wait goroute ", i, " exited")
		}
		close(resultChan)
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}
