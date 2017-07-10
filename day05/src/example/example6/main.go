package main

import "fmt"

type car struct {
	Color string
}

type student struct {
	car
	Name string
	Age int
	Score float32
}

func main() {
	var stu student

	//stu = student{car:car{"white"},Name:"huangwei",Age:20,Score:20}

	stu.Color = "WHITE"
	stu.car.Color = "yellow"

	stu.Name = "huangwei"
	stu.Age = 20
	stu.Score = 50

	fmt.Println(stu);
}
