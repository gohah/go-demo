package main

import (
	"html/template"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age int
}

func Tmpl(writer http.ResponseWriter, request * http.Request) {
	tmpl, err := template.ParseFiles("/Users/victor/learning/go/day10/src/example/template/index.html")

	if err != nil {
		fmt.Println("没找到相应模板",err)
		return
	}

	var pArr []Person

	p1 := Person{"huangwei", 26}
	p2 := Person{"huangwei", 26}
	p3 := Person{"huangwei", 26}
	p4 := Person{"huangwei", 26}
	p5 := Person{"huangwei", 26}

	pArr = append(pArr,p1)
	pArr = append(pArr,p2)
	pArr = append(pArr,p3)
	pArr = append(pArr,p4)
	pArr = append(pArr,p5)

	err = tmpl.Execute(writer,pArr)

	if err != nil {
		fmt.Println("执行模板错误",err)
	}
}

func main() {
	http.HandleFunc("/",Tmpl)
	err := http.ListenAndServe("0.0.0.0:8888",nil)

	if err != nil {
		fmt.Println("监听错误",err)
	}
}
