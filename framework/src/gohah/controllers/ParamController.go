package controllers

import (
	"github.com/astaxie/beego"
	"gohah/models"
	"fmt"
	"context"
)

type ParamController struct {
	beego.Controller
}

func (this *ParamController)Post() {
	stu := &models.Student{}

	if err := this.ParseForm(stu); err != nil {

	}
	fmt.Println(stu)
	this.Ctx.WriteString(stu.Name)
}

func (this *ParamController)Get() {


	this.SetSession("admin","aaaaaa")
	sess := this.GetSession("admin")

	this.Ctx.WriteString(sess.(string))


}
