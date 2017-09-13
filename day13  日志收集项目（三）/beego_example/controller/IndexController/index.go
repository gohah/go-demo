package IndexController

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type IndexController struct {
	beego.Controller
}

func (p *IndexController) Index() {

	logs.Debug("enter index controller")
	p.TplName = "index/index.html"
}
