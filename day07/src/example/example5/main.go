package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

type charCount struct {
	alphaCount int
	digitCount int
	spaceCount int
	otherCount int
}

func main() {

	file,err := os.Open("/Users/victor/test.log")

	if err != nil {
		fmt.Println("not found file: ",err)
	}

	reader := bufio.NewReader(file)

	var count charCount

	for {
		str,err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		fmt.Println(str)
		arr := []rune(str)
		fmt.Println(arr)


		for _,v := range arr {
			fmt.Println(v)
			fmt.Println("===========")
			switch  {
			case v >= 'a' && v <='z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.alphaCount++
			case v == ' ' || v == '\t':
				count.spaceCount++
			case v > '0' && v <= '9':
				count.digitCount++
			default:
				count.otherCount++
			}
		}
	}

	fmt.Println(count.alphaCount)
	fmt.Println(count.spaceCount)
	fmt.Println(count.digitCount)
	fmt.Println(count.otherCount)
}
