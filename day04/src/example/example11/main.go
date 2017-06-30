package main

import "fmt"

func main() {
	//...用法，表示展开切片
	//var a = []int{1,2,3}
	//var b = []int{5,6,7}
	//
	//c := append(a,b...)
	//
	//fmt.Println(c)

	//var a = []int{1,2,3}

	b := make([]int,1)

	b = append(b,1)
	b = append(b,1)
	b = append(b,1)

	//copy(b,a) //拷贝不会扩容

	fmt.Println(b)

}
