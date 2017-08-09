package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)

	var result = 0
	for i := 0; i < len(str); i++ {
		num := int(str[i] - '0')
		result += (num * num * num)
	}

	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("can not convert %s to int\n", str)
		return
	}

	if result == number {
		fmt.Printf("%d is shuixianhuashu\n", number)
	} else {
		fmt.Printf("%d is not shuixianhuashu\n", number)
	}
}
