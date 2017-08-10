package main

import "fmt"

func fab(n int) {
	var a []uint64
	a = make([]uint64, n)

	a[0] = 1
	a[1] = 1

	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	for _, v := range a {
		fmt.Println(v)
	}
}

func testArray() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var a1 = [5]int{1, 2, 3, 4, 5}
	var a2 = [...]int{38, 283, 48, 38, 348, 387, 484}
	var a3 = [...]int{1: 100, 3: 200}
	var a4 = [...]string{1: "hello", 3: "world"}

	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
}

func testArray2() {
	var a [2][5]int = [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	for row, v := range a {
		for col, v1 := range v {
			fmt.Printf("(%d,%d)=%d ", row, col, v1)
		}
		fmt.Println()
	}
}

func main() {
	//testArray()
	testArray2()
	fab(10)
}
