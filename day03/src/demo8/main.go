package main

import(
	"fmt"
)

type Human struct {
	name string
	age int
}

type Student struct {
	Human
	speciality string
}

type Ainimal interface {

}

//断言类型

func main() {
	var a Ainimal;
	a = Student{Human{"gohah",25},"paly"}
	value,ok := a.(Human)
	fmt.Println(value,ok)
}
