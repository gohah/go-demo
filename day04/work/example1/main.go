package main

import "fmt"

func multi() {
	for i := 0; i < 9; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", (i + 1), j+1, (i+1)*(j+1))
		}
		fmt.Println()
	}
}

func main() {
	multi()
}
