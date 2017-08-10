package main

import (
	"fmt"
	"strings"
)

func Adder() func(int) int {
	var x int
	f := func(d int) int {
		x += d
		return x
	}
	return f
}

func makeSuffix(suffix string) func(string) string {
	f := func(name string) string {

		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}

	return f
}

func main() {
	/*
		f := Adder()
		fmt.Println(f(1))
		fmt.Println(f(100))
		fmt.Println(f(1000))
	*/
	f1 := makeSuffix(".bmp")
	fmt.Println(f1("test"))
	fmt.Println(f1("pic"))

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("test"))
	fmt.Println(f2("pic"))
}
