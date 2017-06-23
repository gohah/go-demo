package main

import "fmt"

//设置返回值s
func sum(x,y int)(s int) {
	s = x + y
	return
}

func main() {
	sum := sum(1,2)

	fmt.Println(sum)
}
