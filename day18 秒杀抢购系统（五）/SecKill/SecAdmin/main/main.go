package main

import (
	"github.com/astaxie/beego"
	_ "go_dev/day14/SecKill/SecAdmin/router"
	"fmt"
)

func main() {
	err := initAll()
	if err != nil {
		panic(fmt.Sprintf("init database failed, err:%v", err))
		return
	}
	beego.Run()
}
