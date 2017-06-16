package main

import(
	"fmt"
	"time"
)
//有时候会出现goroutine阻塞的情况，那么我们如何避免整个的程序进入阻塞的情况呢?我们可以利用select来设置 超时，通过如下的方式实现:
func main() {
	c:=make(chan int)
	o:=make(chan bool)

	go func() {
		select {
		case v := <- c:
			fmt.Println(v)
		case <- time.After(5 * time.Second):
			fmt.Println("timeout")
			o <- true
			break
		}

	}()
	<- o
}
