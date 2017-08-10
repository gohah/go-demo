package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	for i := 0; i < 16; i++ {
		go func() {
			for {
				t := time.NewTicker(time.Second)
				select {
				case <-t.C:
					fmt.Println("timeout")
				}
				t.Stop()
			}
		}()
	}

	time.Sleep(time.Second * 100)
}
