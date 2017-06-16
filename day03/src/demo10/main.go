package main

import(
	"fmt"
)

type Person interface {
	say()
	bark()
	eat()
}

type Student interface {
	Person
	study()
}


type StudentA struct {

}

func (a StudentA)say() {

}

func (a StudentA)bark() {

}

func (a StudentA)eat() {

}

func (a StudentA)study() {

}



func main() {
	var s Student;
	s = StudentA{}
	value,ok := s.(Person)
	fmt.Println(value,ok)
}
