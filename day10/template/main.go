package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name  string
	Title string
	age   string
}

func main() {
	t, err := template.ParseFiles("d:/project/src/go_dev/day10/template/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "Mary", age: "31", Title: "我的个人网站"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
