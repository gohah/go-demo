package main

import "fmt"

//函数多值传递
func add(x,y int,z ...int)(s int) {
	s = x +y
	for i:=0; i<len(z); i++ {
		s += z[i]
	}
	return
}

func main() {
	fmt.Println(add(1,2,3,4,5,6,7,8,9,10))
}
