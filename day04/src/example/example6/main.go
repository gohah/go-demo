package main

import "fmt"

func test(a []int) {
	a[0] = 1
}

func main() {

	//var a [5]int

	a := make([]int,5)


	a[0]= 10

	test(a)

	fmt.Println(a)
}
