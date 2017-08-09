package main

import "fmt"

func sum(n int) uint64 {

	var s uint64 = 1
	var sum uint64 = 0
	for i := 1; i <= n; i++ {
		s = s * uint64(i)
		fmt.Printf("%d!=%v\n", i, s)
		sum += s
	}
	return sum
}

func main() {
	var n int

	fmt.Scanf("%d", &n)

	s := sum(n)
	fmt.Println(s)
}
