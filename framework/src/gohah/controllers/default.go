package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "index.html"
}
