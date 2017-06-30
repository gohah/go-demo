package main

import (
	"bufio"
	"os"
	"fmt"
)
//计算字符串中字母数字空格中文字符的个数
func count(str string)(wordCount, spaceCount, numberCount,otherCount int) {
	t := []rune(str)

	for _,v := range t {
		switch  {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			wordCount++
		case v >= '0' && v <= '9':
			numberCount++
		case v==' ':
			spaceCount++
		default:
			otherCount++
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	result,_,err := reader.ReadLine()

	if err == nil {
		fmt.Println("read from console err ",err)
	}

	a,b,c, d := count(string(result))

	fmt.Println(a,b,c,d)
}
