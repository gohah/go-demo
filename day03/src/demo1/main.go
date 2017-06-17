package main

import(
	"fmt"
)

type Person struct {
	name string
	age int
}

func main() {

	//p := Person{"gohah",22}
	//p := Person{name:"gohah",age:23}
	//
	//fmt.Println(p.name)
	//
	//fmt.Println(p.age)

	var str string = "abcdefghijklmnopqrstuvwxyz"

	fmt.Println(str[1:2])
}
