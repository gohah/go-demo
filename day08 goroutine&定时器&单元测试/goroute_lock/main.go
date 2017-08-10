package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m    = make(map[int]uint64)
	lock sync.Mutex
)

type task struct {
	n int
}

func calc(t *task) {
	var sum uint64
	sum = 1
	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}

	fmt.Println(t.n, sum)
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

func main() {
	for i := 0; i < 16; i++ {
		t := &task{n: i}
		go calc(t)
	}

	time.Sleep(10 * time.Second)
	lock.Lock()
	for k, v := range m {
		fmt.Printf("%d! = %v\n", k, v)
	}
	lock.Unlock()
}
