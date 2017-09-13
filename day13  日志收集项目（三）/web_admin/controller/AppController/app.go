package AppController

import (
	"fmt"
	"go_dev/day13/web_admin/model"

	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AppController struct {
	beego.Controller
}

func (p *AppController) AppList() {

	logs.Debug("enter index controller")

	p.Layout = "layout/layout.html"
	appList, err := model.GetAllAppInfo()
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("服务器繁忙")
		p.TplName = "app/error.html"

		logs.Warn("get app list failed, err:%v", err)
		return
	}

	logs.Debug("get app list succ, data:%v", appList)
	p.Data["applist"] = appList

	p.TplName = "app/index.html"
}

func (p *AppController) AppApply() {

	logs.Debug("enter index controller")
	p.Layout = "layout/layout.html"
	p.TplName = "app/apply.html"
}

func (p *AppController) AppCreate() {

	logs.Debug("enter index controller")
	appName := p.GetString("app_name")
	appType := p.GetString("app_type")
	developPath := p.GetString("develop_path")
	ipstr := p.GetString("iplist")

	p.Layout = "layout/layout.html"

	if len(appName) == 0 || len(appType) == 0 || len(developPath) == 0 || len(ipstr) == 0 {
		p.Data["Error"] = fmt.Sprintf("非法参数")
		p.TplName = "app/error.html"

		logs.Warn("invalid parameter")
		return
	}

	appInfo := &model.AppInfo{}
	appInfo.AppName = appName
	appInfo.AppType = appType
	appInfo.DevelopPath = developPath
	appInfo.IP = strings.Split(ipstr, ",")

	err := model.CreateApp(appInfo)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("创建项目失败，数据库繁忙")
		p.TplName = "app/error.html"

		logs.Warn("invalid parameter")
		return
	}

	p.Redirect("/app/list", 302)
}
