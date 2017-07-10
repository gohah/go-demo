package main

import "fmt"

type Student struct {
	Name string
	Age int32
	Score float32
}

func main() {
	//var stu Student
	//
	//stu.Name ="huaxxxx"
	//
	//stu.Age =20
	//
	//stu.Score = 20
	//
	//fmt.Println(stu)
	//fmt.Println(&stu.Name)
	//fmt.Println(&stu.Age)
	//fmt.Println(&stu.Score)

	var stu2 Student = Student{Name:"HUANGWEI",Age:20,Score:30}
	var stu3 *Student = &Student{Name:"HUANGWEI",Age:20,Score:30}

	fmt.Println(stu2.Name,stu3.Name)

}
