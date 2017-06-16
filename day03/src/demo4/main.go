package main

import(
	"fmt"
)

type Human struct {
	name string
	age int
}

func (h Human)say() {
	fmt.Println("human "+h.name+" is say hi")
}

type Person struct {
	Human
	speciality string
}


func (p Person)say() {
	fmt.Println("person "+p.name+" is say hi")
}



func main() {
	p := Person{Human{"gohah",20},"play basketball"}

	p.say()
}
