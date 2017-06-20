package main

import (
	"net/http"
	"log"
)

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("say bye v2..."))
}

//自定义路由Mux，
func main() {
	mux := http.NewServeMux()

	mux.Handle("/",&Myhandler{})

	mux.HandleFunc("/bye",sayBye)

	log.Println("start,v2....")

	log.Fatal(http.ListenAndServe(":4000",mux))
}

type Myhandler struct {

}

func (_ *Myhandler)ServeHTTP(w http.ResponseWriter, r * http.Request) {
	w.Write([]byte("hello,v2....."))
}


