package test

import "fmt"

var Name string = "gohah2"

var Age int = 55


func init() {
	fmt.Println("test init .....")
	fmt.Printf("test Name is %s\n",Name)
	fmt.Printf("test Age is %d\n",Age)
}
