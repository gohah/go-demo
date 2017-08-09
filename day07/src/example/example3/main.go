package main

import (
	"strings"
	"fmt"
)

func main() {
	str := "a b c"

	strSlice := strings.Fields(str)

	fmt.Println(strSlice)
}
