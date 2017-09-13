package LogController

import (
	"fmt"
	"go_dev/day13/web_admin/model"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LogController struct {
	beego.Controller
}

func (p *LogController) LogList() {

	logs.Debug("enter index controller")

	p.Layout = "layout/layout.html"
	logList, err := model.GetAllLogInfo()
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("服务器繁忙")
		p.TplName = "app/error.html"

		logs.Warn("get app list failed, err:%v", err)
		return
	}

	logs.Debug("get app list succ, data:%v", logList)
	p.Data["loglist"] = logList

	p.TplName = "log/index.html"
}

func (p *LogController) LogApply() {

	logs.Debug("enter index controller")
	p.Layout = "layout/layout.html"
	p.TplName = "log/apply.html"
}

func (p *LogController) LogCreate() {

	logs.Debug("enter index controller")
	appName := p.GetString("app_name")
	logPath := p.GetString("log_path")
	topic := p.GetString("topic")

	p.Layout = "layout/layout.html"
	if len(appName) == 0 || len(logPath) == 0 || len(topic) == 0 {
		p.Data["Error"] = fmt.Sprintf("非法参数")
		p.TplName = "log/error.html"

		logs.Warn("invalid parameter")
		return
	}

	logInfo := &model.LogInfo{}
	logInfo.AppName = appName
	logInfo.LogPath = logPath
	logInfo.Topic = topic

	err := model.CreateLog(logInfo)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("创建项目失败，数据库繁忙")
		p.TplName = "log/error.html"

		logs.Warn("invalid parameter")
		return
	}

	iplist, err := model.GetIPInfoByName(appName)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("获取项目ip失败，数据库繁忙")
		p.TplName = "log/error.html"

		logs.Warn("invalid parameter")
		return
	}
	keyFormat := "/oldboy/backend/logagent/config/%s"

	for _, ip := range iplist {
		key := fmt.Sprintf(keyFormat, ip)
		err = model.SetLogConfToEtcd(key, logInfo)
		if err != nil {
			logs.Warn("Set log conf to etcd failed, err:%v", err)
			continue
		}
	}
	p.Redirect("/log/list", 302)
}
