package main

import "fmt"

type Car struct {
	weight int
	name   string
}

func (p *Car) Run() {
	fmt.Println("running")
}

type Bike struct {
	Car
	lunzi int
}

type Train struct {
	Car
}

func (p *Train) String() string {
	str := fmt.Sprintf("name=[%s] weight=[%d]", p.name, p.weight)
	return str
}

func main() {
	var a Bike
	a.weight = 100
	a.name = "bike"
	a.lunzi = 2

	fmt.Println(a)
	a.Run()

	var b Train
	b.weight = 100
	b.name = "train"
	b.Run()

	fmt.Printf("%s", &b)
}
