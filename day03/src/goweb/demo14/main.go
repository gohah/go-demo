package main

import "fmt"

func test(i chan int) {
	i<-11
	i<-1

	i<-1
	i<-1
}

func main() {
	var i chan int;

	i = make(chan int,1)
	go test(i)
	c :=<- i
	d :=<- i
	e :=<- i
	f :=<- i

	fmt.Println(c,d,e,f)
}
