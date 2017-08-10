package main

import "fmt"

type student struct {
	name string
}

func main() {

	var stuChan chan interface{}
	stuChan = make(chan interface{}, 10)

	stu := student{name: "stu01"}

	stuChan <- &stu

	var stu01 interface{}
	stu01 = <-stuChan

	var stu02 *student
	stu02, ok := stu01.(*student)
	if !ok {
		fmt.Println("can not convert")
		return
	}

	fmt.Println(stu02)
}
