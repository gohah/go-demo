package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	str,err := reader.ReadString('\n')

	if err != nil {
		fmt.Println()
	}

	fmt.Println(str)

}
