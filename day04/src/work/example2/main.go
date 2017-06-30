package main

import "fmt"

//完数
func isPrefect(n int)bool {
	//var sum int = 0
	sum := 0
	for i:=1; i<n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return sum == n
}

func proccess(n int) {
	for i:=1; i<n+1; i++ {
		if isPrefect(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	//proccess(1000)
	var n int
	fmt.Scanf("%d",&n)

	proccess(n)
}