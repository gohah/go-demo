package main

import "fmt"

//传递func型类型

type add_func func(x,y int)int

func add(x,y int)int {
	return x + y
}

func operator(op add_func,a int, b int) int {
	return op(a,b)
}

func main() {
	var c add_func = add

	d := operator(c, 100,200)

	fmt.Println(d)
}
