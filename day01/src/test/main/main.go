package main

import(
	"test/mymath"
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

func main() {

	x := mymath.Sqrt(12.0)
	fmt.Println(x)
}


