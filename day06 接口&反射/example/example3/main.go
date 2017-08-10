package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name     string
	Id       string
	Age      int
	sortType int
}

type Book struct {
	Name   string
	Author string
}

type StudentArray []Student

func (p StudentArray) Len() int {
	return len(p)
}

func (p StudentArray) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

func (p StudentArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var stus StudentArray
	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%d", rand.Intn(100)),
			Id:   fmt.Sprintf("110%d", rand.Int()),
			Age:  rand.Intn(100),
		}

		stus = append(stus, stu)
	}

	for _, v := range stus {
		fmt.Println(v)
	}

	fmt.Println("\n\n")
	sort.Sort(stus)
	for _, v := range stus {
		fmt.Println(v)
	}
}
