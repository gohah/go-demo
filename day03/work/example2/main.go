package main

import "fmt"

func isNumber(n int) bool {
	var i, j, k int
	i = n % 10
	j = (n / 10) % 10
	k = (n / 100) % 10

	sum := i*i*i + j*j*j + k*k*k
	return sum == n
}

func main() {
	var n int
	var m int

	fmt.Scanf("%d,%d", &n, &m)

	for i := n; i < m; i++ {

		if isNumber(i) == true {
			fmt.Println(i)
		}
	}
}
