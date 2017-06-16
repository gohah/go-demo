package main

import(
	"fmt"
)

type Student struct {
	name string
	age int
}

//相当于java中的 toString
func (s Student)String() string {
	return "string student"
}

func main() {
	s := Student{"huangwei",22}
	fmt.Println(s)
}