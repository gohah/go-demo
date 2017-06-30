package main

import (
	"errors"
	"fmt"
)

func initConfig() (err error) {
	return errors.New("init config failed")
}

func test() {
	//捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//b := 0
	//
	//a := 10/b

	//

	err := initConfig();
	if err != nil {
		panic(err)
	}

}

func main() {
	/*NEW 初始化值类型，MAKE初始化引用类型*/
	//var a int
	//
	//fmt.Println(a)
	//
	//b := new(int)
	//
	//fmt.Println(b)

	//数组，SLICE append方法, a...代表展开数组或slice
	//var a []int
	//
	//a = append(a,10,20,100)
	//
	//a = append(a, a...)
	//
	//fmt.Println(a)
	for {
		test()
	}

}
