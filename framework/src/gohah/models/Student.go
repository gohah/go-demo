package models

import "fmt"

type Student struct {
	Id int `form:"-"`
	Name string `form:"username"`
	Age string `form:"age"`
	Sex string `form:"sex"`
	Email string
}

func (stu *Student)String() string {
	return fmt.Sprintln(stu.Name,stu.Age,stu.Email,stu.Sex)
}