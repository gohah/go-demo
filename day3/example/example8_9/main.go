package main

import "fmt"

func modify(p *int) {

	fmt.Println(p)
	*p = 1000900
	return
}

func main() {

	var a int = 10
	fmt.Println(&a)

	var p *int
	p = &a

	fmt.Println("the address of p:", &p)
	fmt.Println("the value of p:", p)
	fmt.Println("the value of p point to variable:", *p)

	fmt.Println(*p)
	*p = 100
	fmt.Println(a)

	var b int = 999
	p = &b
	*p = 5

	fmt.Println(a)
	fmt.Println(b)

	modify(&a)
	fmt.Println(a)
}
