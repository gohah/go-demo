package main

import (
	"fmt"
	"time"
)

type Car struct {
	Name string
	Age  int
}

func (c *Car) Set(name string, age int) {
	c.Name = name
	c.Age = age
}

type Car2 struct {
	Name string
}

type Train struct {
	Car
	Car2
	createTime time.Time
	int
}

func (t *Train) Set(age int) {
	t.int = age
}

func main() {
	var train Train
	train.int = 300
	train.Car.Set("huas", 100)

	train.Car.Name = "test"
	fmt.Println(train)
}
