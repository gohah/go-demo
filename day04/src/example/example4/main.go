package main

import "fmt"

//闭包例子1

//20
//120
func main() {
	add := Adder()
	fmt.Println(add(10))
	fmt.Println(add(100))
}

func Adder() func(n int) int {

	var x = 10

	return func(a int) int {
		x += a
		return x
	}
}
