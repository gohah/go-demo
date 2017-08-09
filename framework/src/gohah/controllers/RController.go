package controllers

import "github.com/astaxie/beego"

type RController struct {
	beego.Controller
}

func (this *RController)Get() {
	id := this.Ctx.Input.Param(":id")
	hi := this.Ctx.Input.Param(":hi")
	//path := this.Ctx.Input.Param(":path")
	//ext := this.Ctx.Input.Param(":ext")
	//splat := this.Ctx.Input.Param(":splat")

	this.Ctx.WriteString(id)
	this.Ctx.WriteString(hi)
	//this.Ctx.WriteString(path)
	//this.Ctx.WriteString(ext)
	//this.Ctx.WriteString(splat)
}
