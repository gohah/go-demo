package main

import "fmt"

func main() {
	var a int = 10

	fmt.Println(&a)

	var p *int

	p = &a

	fmt.Println(*p)
}