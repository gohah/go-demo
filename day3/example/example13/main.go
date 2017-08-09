package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func test() {
	return
}

func main() {

	c := add
	sum := c(200, 300)
	fmt.Println(sum)

	str := "hello, world,中国"

	for index, val := range str {
		fmt.Printf("index[%d] val[%c] len[%d]\n", index, val, len([]byte(string(val))))
	}
}
