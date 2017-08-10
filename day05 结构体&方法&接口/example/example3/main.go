package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
	left  *Student
	right *Student
}

func trans(root *Student) {
	if root == nil {
		return
	}
	fmt.Println(root)

	trans(root.left)
	trans(root.right)

}

func main() {
	var root *Student = new(Student)

	root.Name = "stu01"
	root.Age = 18
	root.Score = 100

	var left1 *Student = new(Student)
	left1.Name = "stu02"
	left1.Age = 18
	left1.Score = 100

	root.left = left1

	var right1 *Student = new(Student)
	right1.Name = "stu04"
	right1.Age = 18
	right1.Score = 100

	root.right = right1

	var left2 *Student = new(Student)
	left2.Name = "stu03"
	left2.Age = 18
	left2.Score = 100

	left1.left = left2

	trans(root)
}
