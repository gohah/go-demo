package main

import "fmt"

//9*9乘法表
func multi() {
	for i:=0; i<9; i++ {
		for j:=0; j<=i; j++ {
			fmt.Printf("%d * %d = %d ", (i+1),(j+1),(i+1)*(j+1))
		}
		fmt.Println();
	}
}

func main() {
	multi()
}
