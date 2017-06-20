package main

import(
	"log"
	"net/http"
)

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("say bye..."))
}

//默认路由
func main() {

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello,world"))
	})

	http.HandleFunc("/bye",sayBye)
	log.Println("Start server v1...")
	log.Fatal(http.ListenAndServe(":4000",nil))
}