package main

import (
	"encoding/json"
	"fmt"
)

type student1 struct {
	Name string
	Age  int
}

type student2 struct {
	Name string
	Age  int
}

func main() {
	var m map[string]interface{}
	m = make(map[string]interface{})

	m["stu1"] = student1{
		Name: "001",
		Age:  12,
	}

	m["stu2"] = student2{
		Name: "002",
		Age:  12,
	}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	err = json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range m {
		fmt.Println("key:", k)
		fmt.Println("val:", v)
	}
}
