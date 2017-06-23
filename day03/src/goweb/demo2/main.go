package main

import(
	"fmt"
)

type Human struct {
	name string
	age int
	height float32
	weight float32
}

type Student struct {
	Human
	speciality string
}

func main() {
	s := Student{Human{"gohah",22,170,60},"打篮球"}

	//fmt.Println(s.name)
	//fmt.Println(s.age)
	//fmt.Println(s.height)
	//fmt.Println(s.weight)
	fmt.Println(s.Human.name)
	fmt.Println(s.Human.age)
	fmt.Println(s.Human.height)
	fmt.Println(s.Human.weight)
	fmt.Println(s.speciality)
}
