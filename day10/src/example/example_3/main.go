package main

import (
	"net/http"
	"fmt"
)



func main() {
	url := []string {
		"http://www.baidu.com",
			"http://www.taobao.com",
			"http://www.sina.com.cn",
	}
	for _,v := range url {
		resp, err := http.Head(v)

		if err != nil {

		}

		fmt.Println(resp.StatusCode)
	}
}
