package main

import (
	"reflect"
	"fmt"
)

type student struct {
	name string
	age int
	sex string
}



func main() {
	var s *student = &student{
		name:"vic",
		age:20,
		sex:"男",
	}

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	kd := v.Elem().Kind()

	f := v.Elem().NumMethod()

	fmt.Println(t,v,kd,f)
}
