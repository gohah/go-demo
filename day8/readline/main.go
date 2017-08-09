package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("C:/test.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var line []byte
	for {
		data, prefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		line = append(line, data...)
		if !prefix {
			fmt.Printf("data:%s\n", string(line))
			line = line[:]
		}

	}
}
