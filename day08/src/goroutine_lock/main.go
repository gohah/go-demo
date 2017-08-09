package main

import (
	"fmt"
	"time"

	"sync"
)

var(
	m = make(map[int]uint64)
	lock  sync.Mutex
)

type task struct {
	n int
}


func test(t *task) {

	var sum uint64 = 1

	for i:=1;i<t.n;i++ {
		sum *= uint64(i)

	}
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

func main() {
	for i:=1; i<16; i++ {
		t := &task{n:i}
		go test(t)
	}
	time.Sleep(10*time.Second)
	lock.Lock()
	for k,v := range m {
		fmt.Printf("%d!=%v\n",k,v)
	}
	lock.Unlock()
}
