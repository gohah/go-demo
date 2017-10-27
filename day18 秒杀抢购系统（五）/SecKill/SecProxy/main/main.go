package main

import (
	_ "go_dev/day14/SecKill/SecProxy/router"

	"github.com/astaxie/beego"
)

func main() {

	err := initConfig()
	if err != nil {
		panic(err)
		return
	}

	err = initSec()
	if err != nil {
		panic(err)
		return
	}

	beego.Run()
}
