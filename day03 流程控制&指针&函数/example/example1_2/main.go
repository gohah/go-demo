package main

import (
	"fmt"
	"strings"
)

func urlProcess(url string) string {

	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}

	return url
}

strings.TrimSpace(" sksk ") =>"sksk"
strings.Trim("abbacba", "ab") =>"c"
strings.TrimLeft("a","b")
strings.TrimRight("b","c")
"heheheworld", "he", "wo", 0

strings.Fields("abc cde edk") ["abc", "cde", "edk]

strings.Split("abc,cde,edk", ",") ["abc", "cde", "edk]
strings.Join(["abc", "cde", "edk], ",") "abc,cde,edk"

strings.Replace("str", 3) "strstrstr"
func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}

	return path
}

func main() {
	var (
		url  string
		path string
	)

	fmt.Scanf("%s%s", &url, &path)
	url = urlProcess(url)
	path = pathProcess(path)

	fmt.Println(url)
	fmt.Println(path)
}
