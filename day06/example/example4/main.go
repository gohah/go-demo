package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("read data")
}

func (f *File) Write() {
	fmt.Println("write data")
}

func Test(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {
	var f *File
	var b interface{}
	b = f
	//Test(&f)
	v, ok := b.(ReadWriter)
	fmt.Println(v, ok)
}
