package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("len of args:%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%d]=%s\n", i, v)
	}
}
