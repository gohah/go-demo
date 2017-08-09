package controllers

import (
	"github.com/astaxie/beego"
)

type ObjectController struct {
	beego.Controller
}

func (this *ObjectController)Login() {
	Pmap := this.Ctx.Input.Params()
	for k,v := range Pmap {
		this.Ctx.WriteString(k +"=>"+ v+"\n")
	}
}
