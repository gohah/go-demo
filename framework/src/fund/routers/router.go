package routers

import (
	"fund/controllers"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
)

func init() {
    	beego.Router("/", &controllers.MainController{})

	//固定路由
	beego.Router("/user",&controllers.UserController{})

	//基础路由
	//beego.Get("/get",func(ctx *context.Context){
	//	ctx.Output.Body([]byte("hello get"))
	//})
	//
	//beego.Post("/post",func(ctx *context.Context){
	//	ctx.Output.Body([]byte("hello,post"))
	//})
	//
	//beego.Any("/foo",func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("hello,any"))
	//})
	//
	////正则路由
	////beego.Router("/api/?:id", &controllers.UserController{})
	//
	//beego.Router("/api/:id([0-9]+)",&controllers.UserController{})

}
