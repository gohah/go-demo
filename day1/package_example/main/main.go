package main


import(
	"go_dev/day1/package_example/calc"
	"fmt"
)

func main() {
	sum := calc.Add(100, 300)
	sub := calc.Sub(100, 300)

	fmt.Println("sum=",sum)
	fmt.Println("sub=", sub)
}