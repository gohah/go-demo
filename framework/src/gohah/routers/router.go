package routers

import (
	"gohah/controllers"
	"github.com/astaxie/beego"
	"context"
)

func init() {

	//固定路由
    	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/",&controllers.UserController{})
	beego.Router("/param",&controllers.ParamController{})
	//beego.Router("/admin",&admin.UserController{})

	//正则路由
	//beego.Router("/api/?:id",&controllers.RController{})
	//beego.Router("/api/:id",&controllers.RController{})
	//beego.Router("/api/:id([0-9]+)",&controllers.RController{})
	//beego.Router("/api/:username([\\w]+)",&controllers.RController{})
	//beego.Router("/download/*.*",&controllers.RController{})
	//beego.Router("/download/*",&controllers.RController{})
	//beego.Router("/api/:id:int", &controllers.RController{})
	//beego.Router("/api/:hi:string", &controllers.RController{})
	//beego.Router("/api/cms_:id([0-9]+).html", &controllers.RController{})

	//自定义方法及 RESTful 规则

	//beego.Router("/",&controllers.UserController{},"*:Index")

	//自动匹配
	//beego.AutoRouter(&controllers.ObjectController{})

	//注解路由
	//beego.Include(&controllers.CMSController{})


}
