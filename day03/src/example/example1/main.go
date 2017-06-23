package main

import (
	"strings"
	"fmt"
)
//判断字符串前缀后缀
func urlProccess(url string) string {
	result := strings.HasPrefix(url, "http://")

	if !result {
		url = fmt.Sprintf("http://%s",url)
	}
	return url
}

func pathProccess(url string) string {
	result := strings.HasSuffix(url,"/")

	if !result {
		url = fmt.Sprintf("%s/",url)
	}

	return url
}

func main() {
	var str string
	fmt.Scanf("%s\n",&str)

	fmt.Println(pathProccess(urlProccess(str)))

}
