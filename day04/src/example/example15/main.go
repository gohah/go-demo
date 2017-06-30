package main

import (
	"math/rand"
	"fmt"
	"sync"
)
var lock sync.Mutex
//互斥锁
func testMutex() {
	m := map[int]int{1:1,2:2,3:3,4:4}

	for i:=0; i<2; i++ {
		go func(a map[int]int) {
			lock.Lock()
			a[1] = rand.Int()
			lock.Unlock()
		}(m)
	}

	lock.Lock()
	fmt.Println(m)
	lock.Unlock()
}

func main() {
	testMutex()
}