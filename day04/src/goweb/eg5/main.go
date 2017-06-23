package main

import (
	"net/http"
	"log"
	"html/template"
	"fmt"
)

//GO模板文件

type Person struct {
	Name string
	Age int
}

func testFile(w http.ResponseWriter, r * http.Request) {

	//w.Write([]byte("hello, v4"))

	tmpl,err := template.ParseFiles("eg5/main.tmpl")

	if err != nil {
		fmt.Fprintf(w,"ParseFiles: %v", err)
		return
	}

	err = tmpl.Execute(w,&Person{"gohah",20})

	if err != nil {
		fmt.Fprintf(w,"Execute: %v",err)
		return
	}

}


func main() {
	http.HandleFunc("/",testFile)

	log.Println("start v4.....")
	log.Fatal(http.ListenAndServe(":4000",nil))
}
