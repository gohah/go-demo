package main

import "fmt"

//打印num1 to num2的所有水仙花数
func isNumber(num int) bool {
	i,j,k := num%10,num/10%10,num/100%10

	//i:= num%10
	//j:= (num/10)%10
	//k:= (num/100)%10

	sum := i*i*i + j*j*j + k*k*k

	return sum == num
}
func main() {

	var (
		num1 int
		num2 int
	)

	fmt.Scanf("%d %d\n",&num1, &num2)

	for i:=num1; i<num2; i++ {
		if isNumber(i) == true {
			fmt.Println(i)
		}
	}

}
