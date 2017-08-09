package controllers

import "github.com/astaxie/beego"

type CMSController struct {
	beego.Controller
}

func (this *CMSController)URLMapping() {
	this.Mapping("StaticBlock",this.StaticBlock)
	this.Mapping("AllBlock",this.AllBlock)
}
// @router /staticblock/:key [get]
func (this *CMSController)StaticBlock() {
	this.Ctx.WriteString("staticBlock")
}
// @router /allblock/:key [get]
func (this *CMSController)AllBlock() {
	this.Ctx.WriteString("allBlock")

}
