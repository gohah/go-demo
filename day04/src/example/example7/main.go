package main

import "fmt"

func main() {
	var a [5]int = [5]int{1,2,3}

	var b = [5]int{1,2,3,4,5}

	var c = [...]string{"h","u","a","n","g"}

	var d = [...][2]int{{1,2},{3,4}}

	var e = [...]int{1:2,5:3}

	var f = [...]string{1:"huang",2:"wei"}

	fmt.Println(a,b,c,d,e,f)
}
