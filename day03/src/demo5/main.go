package main

import(
	"fmt"
)

type Human struct {
	name string
	age int
}

func (h Human)eat() {
	str:=fmt.Sprintf("%s is eating",h.name)
	fmt.Println(str)
}

func (h Human)bark() {
	fmt.Printf("%s is barking",h.name)
}

type Student struct {
	Human
	speciality string
}

func (s Student)say() {
	fmt.Printf("%s is saying",s.name)
}

type  men interface {
	eat()
	bark()
}

func main() {
	m := Student{Human{"gohah",25},"play football"}
	m.eat()
}
