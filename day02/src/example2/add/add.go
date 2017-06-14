package add

import (
	_ "example2/test" //_表示导入不用的包，只做初始化
	"fmt"
)

//var Name string = "gohah"
//
//var Age int = 25



//一般采用申明的时候初始化

//*** go是编译型语言，在全局只能申明或申明并初始化，
// " = 或 := "是执行语句，只能在函数中执行，不能独立于函数外

var Name string

var Age int

func Test() {
	//Name = "gohah"
	//Age = 25
}

func init() {
	Name = "gohah"
	Age = 25
	fmt.Println("add init.....")

	fmt.Printf("add Name is %s\n",Name)
	fmt.Printf("add Age is %d\n",Age)
}
