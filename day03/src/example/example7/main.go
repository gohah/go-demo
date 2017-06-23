package main

import (
	"fmt"
	"math/rand"
)
//猜数字游戏
func main() {

	n:= rand.Intn(100)

	for {
		var num int

		fmt.Scanf("%d\n",&num)

		var flag bool = false

		switch {

		case num == n:
			fmt.Println("you are right")
			flag = true
		case num > n:
			fmt.Println("you guess uppper")
		case num <n :
			fmt.Println("you guess lower")

		}

		if flag {
			break
		}
	}

}
