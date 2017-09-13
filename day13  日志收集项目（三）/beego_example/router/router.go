package router

import (
	"go_dev/day13/beego_example/controller/IndexController"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")
}
