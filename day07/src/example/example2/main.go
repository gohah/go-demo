package main

import "fmt"

func main() {
	//str := "hello world gohah"

	var (
		a string
		b string
		c string
	)

	//fmt.Sscanf(str, "%s %s %s",&a, &b, &c)

	//fmt.Scanln(&a,&b,&c)

	fmt.Scanf("%s + %s + %s",&a,&b,&c)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
