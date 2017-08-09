package main

import "fmt"

func main() {
	var sum int = 0
	for i:=1; i<=100; i++ {
		sum = sum + i
	}

	fmt.Println(sum)
}
