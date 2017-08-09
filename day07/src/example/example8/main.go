package main

import (
	"flag" // command line option parser
	"fmt"
)

func main() {

	var test bool
	var str string
	var count int
	flag.BoolVar(&test, "b", false, "print on newline")
	flag.StringVar(&str, "s", "", "print on newline")
	flag.IntVar(&count, "c", 1001, "print on newline")
	flag.Parse()

	fmt.Println(test)
	fmt.Println(str)
	fmt.Println(count)
}
