package main

import(
	"fmt"
	a "example2/add"//a 为包别名
)

func main() {
	a.Test()
	fmt.Println(a.Name,a.Age)
}

//init函数是先于main()函数执行
func init() {
	fmt.Println("main init.....")
}
