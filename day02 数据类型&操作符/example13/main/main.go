package main

import (
	"fmt"
)

func main() {
	var n int16 = 34
	var m int32
	//m = n
	m = int32(n)

	fmt.Printf("m=%d n=%d\n", m, n)
}
