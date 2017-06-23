package main

import "fmt"

//1 to num阶乘和
func sumFactorial(num int) int64 {

	var s int64 = 1

	var sum int64 = 0

	for i:=1; i<num; i++ {
		s = s * int64(i)
		sum += s
	}
	return sum
}

func main() {
	var num int

	num = 100

	sum := sumFactorial(num)

	fmt.Println(sum)
}

