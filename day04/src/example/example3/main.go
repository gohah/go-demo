package main

import "fmt"

func cal(n int) int {
	if n == 1 {
		return 1
	}

	return cal(n -1) * n
}

func main() {
	a := cal(5)

	fmt.Println(a)
}
