package activity

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
//	"fmt"
	"go_dev/day14/SecKill/SecAdmin/model"
	"fmt"
	"net/http"
)


type ActivityController struct {
	beego.Controller
}


func (p *ActivityController)CreateActivity(){
	p.TplName = "activity/create.html"
	p.Layout = "layout/layout.html"
	return
}


func (p *ActivityController) ListActivity(){

	activityModel := model.NewActivityModel()
	activityList, err := activityModel.GetActivityList()
	if err != nil {
		logs.Warn("get activity list failed, err:%v", err)
		return
	}

	p.Data["activity_list"] = activityList
	p.TplName = "activity/list.html"
	p.Layout = "layout/layout.html"
	return
}


func (p *ActivityController) SubmitActivity(){

	activityModel := model.NewActivityModel()
	var activity model.Activity

	p.TplName = "activity/list.html"
	p.Layout = "layout/layout.html"

	var err error
	var Error string = "success"
	defer func(){
		if err != nil {
			p.Data["Error"] = Error
			p.TplName = "activity/error.html"
		}
	}()

	name := p.GetString("activity_name")
	if (len(name) == 0) {
		Error = "活动的名字不能为空"
		err = fmt.Errorf("activity name can not null")
		return
	}

	productId, err := p.GetInt("product_id")
	if (err != nil) {
		
		err = fmt.Errorf("商品id 非法, err:%v", err)
		Error = err.Error()
		return
	}
	
	startTime, err := p.GetInt64("start_time")
	if (err != nil) {
		err = fmt.Errorf("开始时间 非法, err:%v", err)
		Error = err.Error()
		return
	}

	endTime, err := p.GetInt64("end_time")
	if (err != nil) {
		err = fmt.Errorf("结束时间 非法, err:%v", err)
		Error = err.Error()
		return
	}

	total, err := p.GetInt("total")
	if (err != nil) {
		err = fmt.Errorf("商品数量 非法, err:%v", err)
		Error = err.Error()
		return
	}

	speed, err := p.GetInt("speed")
	if (err != nil) {
		err = fmt.Errorf("商品速度 非法, err:%v", err)
		Error = err.Error()
		return
	}

	limit, err := p.GetInt("buy_limit")
	if (err != nil) {
		err = fmt.Errorf("购买限制 非法, err:%v", err)
		Error = err.Error()
		return
	}

	activity.ActivityName = name
	activity.ProductId = productId
	activity.StartTime = startTime
	activity.EndTime = endTime
	activity.Total = total
	activity.Speed = speed
	activity.BuyLimit = limit

	err = activityModel.CreateActivity(&activity)
	if err != nil {
		err = fmt.Errorf("创建活动失败, err:%v", err)
		Error = err.Error()
		return
	}

	p.Redirect("/activity/list", http.StatusMovedPermanently)
	return
}