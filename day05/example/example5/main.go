package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"student_name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {
	var stu Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 80,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode stu failed, err:", err)
		return
	}

	fmt.Println(string(data))
}
