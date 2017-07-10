package controllers

import (
	"github.com/astaxie/beego"

	"strconv"
)

type UserController struct {
	beego.Controller
}

func (this *UserController)Get() {
	val,_ := this.GetInt("name")
	this.Ctx.WriteString(strconv.Itoa(val))
}