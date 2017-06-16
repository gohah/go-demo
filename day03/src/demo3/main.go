package main

import(
	"fmt"
)

type Rectangle struct {
	width,height float32
}

func (r Rectangle) area() float32 {
	return r.width * r.height
}

func main() {
	r := Rectangle{12,12}

	fmt.Println(r.area())
}
