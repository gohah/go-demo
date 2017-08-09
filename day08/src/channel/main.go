package main

import "fmt"

type student struct {
	name string
}

func main() {
	//var a chan int
	//
	//
	//a = make(chan int,10)
	//
	//a <- 10

	//var b chan map[string]string
	//
	//b = make(chan map[string]string)
	//
	//m := make(map[string]string)
	//
	//m["a"] = "a"
	//m["b"] = "a"
	//m["c"] = "a"
	//m["d"] = "a"
	//
	//b <- m

	//var stu chan *student
	//
	//stu = make(chan *student,10)
	//
	//stu <- &student{name:"huangwei"}
	//
	////var stu2 *student
	//
	////<- stu
	//
	//stu2 := <- stu
	//
	//fmt.Println(stu2)

	var b chan interface{}

	b = make(chan interface{},10)

	b <- 1
	b <-"aaa"


	fmt.Println((<-b).(int));
	fmt.Println((<-b).(int));

}
