package main

import "fmt"

type integer int

type student struct {
	Age int
}

type stu student
//给类型取别名
func main() {
	//var a integer = 1000
	//
	//var b int = 100
	//
	//b = int(a)

	var a student = student{20}

	var b stu = stu{30}

	b = stu(a)

	fmt.Println(a,b)
}
