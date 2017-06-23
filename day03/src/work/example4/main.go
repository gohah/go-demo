package main

import (
	"fmt"
	"strconv"
)

//数字字符串转换成数字
func sumFactorial(strnum string) int{
	var result int
	for i:=0; i<len(strnum); i++ {
		v := strnum[i] - '0'
		result += int(v*v*v)
	}
	return result
}

func main() {

	var str string

	fmt.Scanf("%s",&str)

	sum := sumFactorial(str)

	res,_ := strconv.Atoi(str)

	fmt.Println(res,sum)
}


