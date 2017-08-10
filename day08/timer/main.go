package main

import (
	"runtime"
	"time"
)

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	for i := 0; i < 10240; i++ {
		go func() {
			for {
				select {
				case <-time.After(time.Nanosecond):
					//fmt.Println("after")
				}
			}
		}()
	}

	time.Sleep(time.Second * 10000)
}
