package main

import "fmt"

var (
	result = func(a1 int, b1 int) int {
		return a1 + b1
	}
)

func test(a, b int) int {
	result := func(a1 int, b1 int) int {
		return a1 + b1
	}

	return result(a, b)
}

func main() {
	fmt.Println(result(100, 200))

	var i int = 0
	defer fmt.Println(i)
	defer fmt.Println("second")

	i = 10
	fmt.Println(i)
}
