package main

import(
	"fmt"
)
func SumEqN(n int) {

	for x:=0; x<=n; x++ {
		for y:=0; y<=n; y++ {
			if x + y == n {
				fmt.Printf("%d + %d = %d",x,y,n)
				fmt.Println()
			}
		}
	}
}
func main() {
	SumEqN(5)
}
