package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	var m int

	fmt.Scanf("%d%d%s", &n, &m)
	for i := n; i < m; i++ {
		if isPrime(i) == true {
			fmt.Printf("%d\n", i)
			continue
		}
	}
}
