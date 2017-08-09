package main

import "fmt"

func main() {

	var link Link
	for i := 0; i < 10; i++ {
		//intLink.InsertHead(i)
		link.InsertTail(fmt.Sprintf("str %d", i))
	}

	link.Trans()
}
