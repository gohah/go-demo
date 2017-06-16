package main

import(
	"fmt"
)

type Element interface {

}

type List []Element


type Person struct {
	name string
	age int
}

func (p Person)String()string {
	return fmt.Sprintf("%s age is %d",p.name,p.age)
}

func main() {
	list := make(List,3)

	list[0] = 3

	list[1] = "hello"

	list[2] = Person{"gohah",25}

	for index, element := range list {

		//element.(type)只能在switch内使用
		switch value := element.(type) {
		case int:
			fmt.Println("int" + string(index) + string(value))
		case string:
			fmt.Println("string"+ string(index) + string(value))
		case Person:
			fmt.Println("Student")
		}
	}
}
