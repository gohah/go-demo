package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string
	var err error
	var num1 int

	fmt.Scanf("%s",&num)

	num1,err = strconv.Atoi(num)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num1)
	}


}
