package main

import "fmt"

/*num1 to num2 所有素数*/
func isPrime(num int) bool {
	for i:=2; i<num; i++ {
		if num %i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num1 int;

	var num2 int;

	fmt.Println("Please input two number: <num1> <num2>")

	fmt.Scanf("%d %d",&num1, &num2)

	for i:=num1; i<num2; i++ {
		if isPrime(i) == true {
			fmt.Printf("%d\n",i)
			continue
		}
	}

}