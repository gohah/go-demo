package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (this *UserController)Get() {
	name := this.Input().Get("name")
	name2,_ := strconv.Atoi(name)
	this.Ctx.WriteString(string(name2))
}



func (this *UserController)Index() {

}

