package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	score float32
}

func main() {
	var stu Student

	stu.Age = 18
	stu.Name = "hua"
	stu.score = 80

	var stu1 *Student = &Student{
		Age:  20,
		Name: "hua",
	}

	var stu3 = Student{
		Age:  20,
		Name: "hua",
	}
	fmt.Println(stu1.Name)
	fmt.Println(stu3)
	fmt.Printf("Name:%p\n", &stu.Name)
	fmt.Printf("Age: %p\n", &stu.Age)
	fmt.Printf("score:%p\n", &stu.score)
}
