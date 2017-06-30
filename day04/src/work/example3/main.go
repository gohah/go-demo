package main

import "fmt"
//回文1221
func process(str string)bool {
	for i:=0; i<len(str);i++ {
		if i == len(str)/2 {
			break;
		}

		last := len(str) -i -1;

		if str[last] != str[i] {
			return false
		}
	}

	return true
}

func main() {
	var str string

	fmt.Scanf("%s",&str)

	if process(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}