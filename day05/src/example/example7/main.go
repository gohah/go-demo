package main

import "fmt"

type student struct {
	Name string
	Age int
	Score float32

	sex int
}
//值传递
//func (s student)init(name string , age int, score float32) {
//	s.Score=score
//	s.Name = name
//	s.Age = age
//	s.sex = 1
//
//	fmt.Println(s)
//}
//引用传递
func (s *student)init(name string , age int, score float32) {
	s.Score=score
	s.Name = name
	s.Age = age
	s.sex = 1

	fmt.Println(*s)
}

func main() {
	var stu student

	stu.init("huangwei",12,50)

	fmt.Println(stu)
}
