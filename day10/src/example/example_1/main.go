package main

import (
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe("0.0.0.0:8888",nil)

	if err != nil {
		fmt.Println("http listen failed")
	}
}
