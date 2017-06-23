package main

import "fmt"

func main() {

	//var i int = 100
	//for i > 10 {
	//	fmt.Println(i)
	//}



	//str := "hello, world"
	//
	//for i,v := range str {
	//	fmt.Println(i,v)
	//}


	label1:
		for i:=0; i<5;i++ {
			for j:=0; j<5; j++ {
				if j ==2 {
					continue label1
				}
				fmt.Println(i,j)
			}
		}
}
