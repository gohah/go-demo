package main

import "fmt"

func main() {

	var i int = 0

	label2:
		i++
		if i == 5 {
			return
		}
		fmt.Println(i)
	goto label2
}
