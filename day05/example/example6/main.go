package main

import (
	"fmt"
)

type Cart1 struct {
	name string
	age  int
}

type Cart2 struct {
	name string
	age  int
}

type Train struct {
	Cart1
	Cart2
}

func main() {
	var t Train

	t.Cart1.name = "train"
	t.Cart1.age = 100

	fmt.Println(t)
}
