package main

import "fmt"

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c", str[strLen-i-1])
	}
	return result
}

func reverse1(str string) string {
	var result []byte
	tmp := []byte(str)
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
	}
	return string(result)
}

func main() {
	var str1 = "hello"
	str2 := "world"

	//str3 := str1 + " " + str2
	str3 := fmt.Sprintf("%s %s", str1, str2)
	n := len(str3)

	fmt.Println(str3)

	fmt.Printf("len(str3)=%d\n", n)

	substr := str3[0:5]
	fmt.Println(substr)

	substr = str3[6:]
	fmt.Println(substr)

	result := reverse(str3)
	fmt.Println(result)

	result = reverse1(result)
	fmt.Println(result)
}
