package main

import "fmt"

//defer延迟执行，会把数据压入栈，先进后出，函数return或完成时才执行

func test() {
	i :=0

	defer fmt.Println(i)

	i++

	return
}

func test2() {
	for i:=0; i<10; i++ {
		defer fmt.Println(i)
	}
	return
}

func main() {
	test()
	fmt.Println("******")
	test2()
}