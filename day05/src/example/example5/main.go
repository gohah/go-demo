package main

import (
	"encoding/json"
	"fmt"
)
//接口体TAG
type student struct {
	Name string `json:"name"`

	Age int `json:"age"`
}


func main() {
	var stu student = student{
		Name:"huangwei",
		Age:30,
	}

	data, err := json.Marshal(stu)

	fmt.Println(string(data),err)
}
