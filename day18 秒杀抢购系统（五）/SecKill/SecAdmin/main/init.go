package main

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go_dev/day14/SecKill/SecAdmin/model"
	etcd_client "github.com/coreos/etcd/clientv3"
	"fmt"
	"time"
)

var Db *sqlx.DB
var EtcdClient *etcd_client.Client

func initDb() (err error) {

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", AppConf.mysqlConf.UserName, AppConf.mysqlConf.Passwd,
				AppConf.mysqlConf.Host, AppConf.mysqlConf.Port, AppConf.mysqlConf.Database)
	database, err := sqlx.Open("mysql", dns)
	if err != nil {
		logs.Error("open mysql failed, err:%v ", err)
		return
	}
	Db = database
	logs.Debug("connect to mysql succ")
	return
}

func initEtcd()(err error) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{AppConf.etcdConf.Addr},
		DialTimeout: time.Duration(AppConf.etcdConf.Timeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}

	EtcdClient = cli
	logs.Debug("init etcd succ")
	return
}

func initAll() (err error) {
	err = initConfig()
	if err != nil {
		logs.Warn("init config failed, err:%v", err)
		return
	}

	err = initDb()
	if err != nil {
		logs.Warn("init Db failed, err:%v", err)
		return
	}

	err = initEtcd()
	if err != nil {
		logs.Warn("init etcd failed, err:%v", err)
		return
	}

	err = model.Init(Db, EtcdClient, AppConf.etcdConf.EtcdKeyPrefix, AppConf.etcdConf.ProductKey)
	if err != nil {
		logs.Warn("init model failed, err:%v", err)
		return
	}
	return
}
