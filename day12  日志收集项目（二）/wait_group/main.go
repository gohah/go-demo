package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calc(&wg, i)
	}

	wg.Wait()
	fmt.Println("all goroutine finish")
}
func calc(w *sync.WaitGroup, i int) {

	fmt.Println("calc:", i)
	time.Sleep(time.Second)
	w.Done()
}
