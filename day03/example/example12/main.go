package main

import "fmt"

func Print(n int) {

	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
}

func main() {
	Print(6)
}
