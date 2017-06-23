package main

import "fmt"

func main() {

	var num int
	num = 5

	switch num {
	case 0,1,2,3,4,5,6,7,8,9:
		fmt.Println(0)
		fallthrough
	case 10:
		fmt.Println(10)
	case 20:
		fmt.Println(20)
	default:
		fmt.Println("default")
	}
}
