package main

import (
	"strings"
	"fmt"
)
//闭包例子
//.jpg
//.bmp
func makeSuffixFunc(suffix string) func(string)string {
	return func(name string) string {
		if !strings.HasSuffix(name,suffix) {
			return name + suffix
		}

		return name
	}
}

func main() {
	func1 := makeSuffixFunc(".jpg")

	func2 := makeSuffixFunc(".bmp")

	fmt.Println(func1("test1"))
	fmt.Println(func1("demo1"))
	fmt.Println(func2("test2"))
	fmt.Println(func2("demo2"))


}
