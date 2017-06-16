package main

import(
	"fmt"
	"net/http"
	"strings"
	"log"
)

//Go语言里面提供了一个完善的net/http包，通过http包可以很 方便的就搭建起来一个可以运行的web服务。同时使用这个包能很简单地对web的路由，静态文件，模版，cookie等数 据进行设置和操作。


func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的 fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息 fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, "")) }
	fmt.Fprintf(w, "Hello web server!") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayHelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//上面这个代码，我们build之后，然后执行web.exe,这个时候其实已经在9090端口监听http链接请求了。 在浏览器输入http://localhost:9090
//可以看到浏览器页面输出了Hello astaxie! 可以换一个地址试试:http://localhost:9090/?url_long=111&url_long=222 看看浏览器输出的是什么，服务器输出的是什么?
