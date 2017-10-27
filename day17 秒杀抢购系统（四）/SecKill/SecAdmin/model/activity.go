package model

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
	"time"
	//"errors"
	"fmt"
	"encoding/json"
	"context"
	"strings"
)

const (
	ActivityStatusNormal = 0
	ActivityStatusDisable = 1
	ActivityStatusExpire = 2
)

type Activity struct {
	ActivityId int `db:"id"`
	ActivityName string `db:"name"`
	ProductId int `db:"product_id"`
	StartTime int64 `db:"start_time"`
	EndTime int64 `db:"end_time"`
	Total int `db:"total"`
	Status int `db:"status"`

	StartTimeStr string
	EndTimeStr string
	StatusStr string
	Speed int `db:"sec_speed"`
	BuyLimit int `db:"buy_limit"`
}

type SecProductInfoConf struct {
	ProductId         int
	StartTime         int64
	EndTime           int64
	Status            int
	Total             int
	Left              int
	OnePersonBuyLimit int
	BuyRate           float64
	//每秒最多能卖多少个
	SoldMaxLimit int
}

type ActivityModel struct {

}

func NewActivityModel() *ActivityModel{
	return &ActivityModel{}
}

func (p *ActivityModel) GetActivityList() (activityList []*Activity, err error) {
	sql := "select id, name, product_id, start_time, end_time, total, status, sec_speed, buy_limit from activity order by id desc"
	err = Db.Select(&activityList, sql)
	if err != nil {
		logs.Error("select activity from database failed, err:%v", err)
		return
	}

	for _, v := range activityList {
		t := time.Unix(v.StartTime, 0)
		v.StartTimeStr = t.Format("2006-01-02 15:04:05")

		t = time.Unix(v.EndTime, 0)
		v.EndTimeStr = t.Format("2006-01-02 15:04:05")

		now := time.Now().Unix()

		if now > v.EndTime {
			v.StatusStr = "已结束"
			continue
		}

		if v.Status == ActivityStatusNormal {
			v.StatusStr = "正常"
		} else if v.Status == ActivityStatusDisable {
			v.StatusStr = "已禁用"
		}
	}
	logs.Debug("get activity succ,  activity list is[%v]", activityList)
	return
}

func (p *ActivityModel) ProductValid(productId int, total int)(valid bool, err error) {
	sql := "select id, name, total, status from product where id=?"
	var productList[]*Product
	err = Db.Select(&productList, sql, productId)
	if err != nil {
		logs.Warn("select product failed, err:%v", err)
		return
	}

	if len(productList) == 0 {
		err = fmt.Errorf("product[%v] is not exists", productId)
		return
	}

	if total > productList[0].Total {
		err = fmt.Errorf("product[%v] 的数量非法", productId)
		return
	}

	valid = true
	return
}

func (p *ActivityModel) CreateActivity(activity *Activity) (err error) {
	
	valid, err := p.ProductValid(activity.ProductId, activity.Total)
	if err != nil {
		logs.Error("product exists failed, err:%v", err)
		return
	}

	if !valid {
		err = fmt.Errorf("product id[%v] err:%v", activity.ProductId, err)
		logs.Error(err)
		return
	}

	if activity.StartTime <= 0 || activity.EndTime <= 0 {
		err = fmt.Errorf("invalid start[%v]|end[%v] time", activity.StartTime, activity.EndTime)
		logs.Error(err)
		return
	}

	if activity.EndTime <= activity.StartTime {
		err = fmt.Errorf("start[%v] is greate then end[%v] time", activity.StartTime, activity.EndTime)
		logs.Error(err)
		return
	}

	now := time.Now().Unix()
	if activity.EndTime <= now || activity.StartTime <= now {
		err = fmt.Errorf("start[%v]|end[%v] time is less then now[%v]", activity.StartTime, activity.EndTime, now)
		logs.Error(err)
		return
	}
	
	sql := "insert into activity(name, product_id, start_time, end_time, total, sec_speed, buy_limit)values(?,?,?,?,?,?,?)"
	_, err = Db.Exec(sql, activity.ActivityName, activity.ProductId, 
		activity.StartTime, activity.EndTime, activity.Total, activity.Speed, activity.BuyLimit)
	if err != nil {
		logs.Warn("select from mysql failed, err:%v sql:%v", err, sql)
		return
	}
	
	logs.Debug("insert into database succ")
	err = p.SyncToEtcd(activity)
	if err != nil {
		logs.Warn("sync to etcd failed, err:%v data:%v", err, activity)
		return
	}
	return
}

func (p *ActivityModel) SyncToEtcd(activity *Activity) (err error) {

	if strings.HasSuffix(EtcdPrefix, "/") == false {
		EtcdPrefix = EtcdPrefix + "/"
	}

	etcdKey  := fmt.Sprintf("%s%s", EtcdPrefix, EtcdProductKey)
	secProductInfoList, err := loadProductFromEtcd(etcdKey)

	var secProductInfo SecProductInfoConf
	secProductInfo.EndTime =  activity.EndTime
	secProductInfo.OnePersonBuyLimit = activity.BuyLimit
	secProductInfo.ProductId = activity.ProductId
	secProductInfo.SoldMaxLimit = activity.Speed
	secProductInfo.StartTime = activity.StartTime
	secProductInfo.Status = activity.Status
	secProductInfo.Total = activity.Total

	secProductInfoList = append(secProductInfoList, secProductInfo)

	data, err := json.Marshal(secProductInfoList)
	if err != nil {
		logs.Error("json marshal failed, err:%v", err)
		return
	}

	_, err = EtcdClient.Put(context.Background(), etcdKey, string(data))
	if err != nil {
		logs.Error("put to etcd failed, err:%v, data[%v]", err, string(data))
		return
	}
	return
}



func loadProductFromEtcd(key string) (secProductInfo []SecProductInfoConf, err error) {

	logs.Debug("start get from etcd succ")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := EtcdClient.Get(ctx, key)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err:%v", key, err)
		return
	}

	logs.Debug("get from etcd succ, resp:%v", resp)
	for k, v := range resp.Kvs {
		logs.Debug("key[%v] valud[%v]", k, v)
		err = json.Unmarshal(v.Value, &secProductInfo)
		if err != nil {
			logs.Error("Unmarshal sec product info failed, err:%v", err)
			return
		}

		logs.Debug("sec info conf is [%v]", secProductInfo)
	}

/*
	updateSecProductInfo(conf, secProductInfo)
	logs.Debug("update product info succ, data:%v", secProductInfo)

	initSecProductWatcher(conf)

	logs.Debug("init etcd watcher succ")
	*/
	return
}