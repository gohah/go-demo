package main

import "fmt"

func main() {
	var a = [...]int{1,2,3,4,5}

	s := a[1:3]

	fmt.Println(a,s,cap(s))

	s[1]= 100

	fmt.Println(a,s,cap(s))

	s = append(s,1000)
	s = append(s,1000)
	s = append(s,1000)
	s = append(s,1000)
	s = append(s,1000)
	s = append(s,1000)
	s = append(s,1000)
	s = append(s,1000)
	s = append(s,1000)
	s = append(s,1000)
	s = append(s,1000)

	s[1] = 10000000

	fmt.Println(a,s,cap(s))

}
