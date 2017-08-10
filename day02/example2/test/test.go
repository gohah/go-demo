package test


import(
	"fmt"
)


var Name string = "this is in test package"
var Age int = 1000

func init() {
	fmt.Println("this is a test")
	fmt.Println("test.package.Name=", Name)
	fmt.Println("test.package.age=", Age)

	Age = 10
	fmt.Println("test.package.age=", Age)
}