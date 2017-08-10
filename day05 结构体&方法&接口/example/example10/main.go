package main

import "fmt"

type People struct {
	name string
	age  int
}

type Test interface {
	Print()
	Sleep()
}

type Student struct {
	name  string
	age   int
	score int
}

func (p Student) Print() {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	fmt.Println("score:", p.score)
}

func (p Student) Sleep() {
	fmt.Println("student sleep")
}

func (people People) Print() {
	fmt.Println("name:", people.name)
	fmt.Println("age:", people.age)
}

func (p People) Sleep() {
	fmt.Println("people sleep")
}

func main() {

	var t Test
	fmt.Println(t)
	//t.Print()

	var stu Student = Student{
		name:  "stu1",
		age:   20,
		score: 200,
	}

	t = stu
	t.Print()
	t.Sleep()

	var people People = People{
		name: "people",
		age:  100,
	}

	t = people
	t.Print()
	t.Sleep()

	fmt.Println("t:", t)
}
