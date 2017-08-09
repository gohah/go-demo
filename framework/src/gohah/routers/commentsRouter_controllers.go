package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gohah/controllers:CMSController"] = append(beego.GlobalControllerRouter["gohah/controllers:CMSController"],
		beego.ControllerComments{
			Method: "AllBlock",
			Router: `/allblock/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gohah/controllers:CMSController"] = append(beego.GlobalControllerRouter["gohah/controllers:CMSController"],
		beego.ControllerComments{
			Method: "StaticBlock",
			Router: `/staticblock/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
