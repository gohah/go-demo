package main

import "fmt"

func test() {
	//申明变量
	a := new([]int)

	fmt.Println(a)

	//初始化
	*a = make([]int,10)

	b := make([]int,10)

	fmt.Println(*a,b)
}

func main() {
	test()
}
