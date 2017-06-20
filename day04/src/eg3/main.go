package main

import (
	"net/http"
	"log"
	"time"
)
//自定义server
func main() {

	server := &http.Server{
		Addr: ":4000",
		WriteTimeout:2*time.Second,
	}

	mux := http.NewServeMux()

	mux.Handle("/",&myHandler{})

	server.Handler = mux

	log.Println("start v3 ....")

	log.Fatal(http.ListenAndServe(":4000",mux))
}


type myHandler struct {

}

func (_ *myHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, v3 ...."))
}
