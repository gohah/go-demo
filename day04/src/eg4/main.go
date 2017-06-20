package main

import (
	"html/template"
	"log"
	"os"
)
//GO模板基本用法
func main() {
	tmpl,err := template.New("go-web").Parse("hello,world! {{.}} My NAME IS GOHAH")

	if(err != nil) {
		log.Fatalf("Parse: %v",err)
	}

	err = tmpl.Execute(os.Stdout,"*****")

	if(err != nil) {
		log.Fatalf("Execute:%v",err)
	}
}
