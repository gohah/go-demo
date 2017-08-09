package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	res, err := http.Get("http://www.liechengcf.com")

	if err != nil {
		fmt.Println("get err :", err)
		return
	}


	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("read err :", err)
		return
	}

	fmt.Println(string(data))
}
