package controller

import (
	"fmt"
	"go_dev/day14/SecKill/SecProxy/service"
	"strings"
	"time"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type SkillController struct {
	beego.Controller
}

func (p *SkillController) SecKill() {

	productId, err := p.GetInt("product_id")
	result := make(map[string]interface{})

	result["code"] = 0
	result["message"] = "success"

	defer func() {
		p.Data["json"] = result
		p.ServeJSON()
	}()

	if err != nil {
		result["code"] = 1001
		result["message"] = "invalid product_id"
		return
	}

	source := p.GetString("src")
	authcode := p.GetString("authcode")
	secTime := p.GetString("time")
	nance := p.GetString("nance")

	secRequest := service.NewSecRequest()
	secRequest.AuthCode = authcode
	secRequest.Nance = nance
	secRequest.ProductId = productId
	secRequest.SecTime = secTime
	secRequest.Source = source
	secRequest.UserAuthSign = p.Ctx.GetCookie("userAuthSign")
	secRequest.UserId, _ = strconv.Atoi(p.Ctx.GetCookie("userId"))
	secRequest.AccessTime = time.Now()
	if len(p.Ctx.Request.RemoteAddr) > 0 {
		secRequest.ClientAddr = strings.Split(p.Ctx.Request.RemoteAddr, ":")[0]
	}

	secRequest.ClientRefence = p.Ctx.Request.Referer()
	secRequest.CloseNotify = p.Ctx.ResponseWriter.CloseNotify()

	logs.Debug("client request:[%v]", secRequest)
	if err != nil {
		result["code"] = service.ErrInvalidRequest
		result["message"] = fmt.Sprintf("invalid cookie:userId")
		return
	}

	data, code, err := service.SecKill(secRequest)
	if err != nil {
		result["code"] = code
		result["message"] = err.Error()
		return
	}

	result["data"] = data
	result["code"] = code

	return
}

func (p *SkillController) SecInfo() {

	productId, err := p.GetInt("product_id")
	result := make(map[string]interface{})

	result["code"] = 0
	result["message"] = "success"

	defer func() {
		p.Data["json"] = result
		p.ServeJSON()
	}()

	if err != nil {
		data, code, err := service.SecInfoList()
		if err != nil {
			result["code"] = code
			result["message"] = err.Error()

			logs.Error("invalid request, get product_id failed, err:%v", err)
			return
		}

		result["code"] = code
		result["data"] = data
	} else {

		data, code, err := service.SecInfo(productId)
		if err != nil {
			result["code"] = code
			result["message"] = err.Error()

			logs.Error("invalid request, get product_id failed, err:%v", err)
			return
		}

		result["code"] = code
		result["data"] = data
	}
}
