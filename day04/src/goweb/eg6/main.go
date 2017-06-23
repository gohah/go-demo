package main

import (
	"net/http"
	"html/template"
	"fmt"
	"log"
	"strconv"
)

//go模板标签使用

func testTemplate(w http.ResponseWriter, r * http.Request) {

	tmpl, err := template.ParseFiles("eg6/main.html")

	if err != nil {
		fmt.Fprintf(w, "ParseFiles: %v", err)
		return
	}

	score := r.FormValue("score")

	scored,_ := strconv.Atoi(score)

	err = tmpl.Execute(w,map[string]interface{}{
		"score":scored,
		"request":r,
		"arr":[]int{1,2,3,4,5},
	})

	if err != nil {
		fmt.Fprintf(w, "Execute: %v",err)
		return
	}
}


func main() {
	http.HandleFunc("/", testTemplate)

	log.Println("start v4...")

	log.Fatal(http.ListenAndServe(":4000",nil))
}
