package main

import (
	"fmt"
	"go_dev/day14/SecKill/SecLayer/service"
	"strings"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

var (
	appConfig *service.SecLayerConf
)

func initConfig(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	//读取日志库配置
	appConfig = &service.SecLayerConf{}
	appConfig.LogLevel = conf.String("logs::log_level")
	if len(appConfig.LogLevel) == 0 {
		appConfig.LogLevel = "debug"
	}

	appConfig.LogPath = conf.String("logs::log_path")
	if len(appConfig.LogPath) == 0 {
		appConfig.LogPath = "./logs"
	}

	//读取redis相关的配置
	appConfig.Proxy2LayerRedis.RedisAddr = conf.String("redis::redis_proxy2layer_addr")
	if len(appConfig.Proxy2LayerRedis.RedisAddr) == 0 {
		logs.Error("read redis::redis_proxy2layer_addr failed")
		err = fmt.Errorf("read redis::redis_proxy2layer_addr failed")
		return
	}

	appConfig.Proxy2LayerRedis.RedisQueueName = conf.String("redis::resis_proxy2layer_queue_name")
	if len(appConfig.Proxy2LayerRedis.RedisQueueName) == 0 {
		logs.Error("read redis::resis_proxy2layer_queue_name failed")
		err = fmt.Errorf("read redis::resis_proxy2layer_queue_name failed")
		return
	}

	appConfig.Proxy2LayerRedis.RedisMaxIdle, err = conf.Int("redis::redis_proxy2layer_idle")
	if err != nil {
		logs.Error("read redis::redis_proxy2layer_idle failed, err:%v", err)
		return
	}

	appConfig.Proxy2LayerRedis.RedisIdleTimeout, err = conf.Int("redis::redis_proxy2layer_idle_timeout")
	if err != nil {
		logs.Error("read redis::redis_proxy2layer_idle_timeout failed, err:%v", err)
		return
	}

	appConfig.Proxy2LayerRedis.RedisMaxActive, err = conf.Int("redis::redis_proxy2layer_active")
	if err != nil {
		logs.Error("read redis::redis_proxy2layer_active failed, err:%v", err)
		return
	}

	//读取redis layer2proxy相关的配置
	appConfig.Layer2ProxyRedis.RedisAddr = conf.String("redis::redis_layer2proxy_addr")
	if len(appConfig.Proxy2LayerRedis.RedisAddr) == 0 {
		logs.Error("read redis::redis_layer2proxy_addr failed")
		err = fmt.Errorf("read redis::redis_layer2proxy_addr failed")
		return
	}

	appConfig.Layer2ProxyRedis.RedisQueueName = conf.String("redis::redis_layer2proxy_queue_name")
	if len(appConfig.Proxy2LayerRedis.RedisQueueName) == 0 {
		logs.Error("read redis::redis_layer2proxy_queue_name failed")
		err = fmt.Errorf("read redis::redis_layer2proxy_queue_name failed")
		return
	}

	appConfig.Layer2ProxyRedis.RedisMaxIdle, err = conf.Int("redis::redis_layer2proxy_idle")
	if err != nil {
		logs.Error("read redis::redis_layer2proxy_idle failed, err:%v", err)
		return
	}

	appConfig.Layer2ProxyRedis.RedisIdleTimeout, err = conf.Int("redis::redis_layer2proxy_idle_timeout")
	if err != nil {
		logs.Error("read redis::redis_layer2proxy_idle_timeout failed, err:%v", err)
		return
	}

	appConfig.Layer2ProxyRedis.RedisMaxActive, err = conf.Int("redis::redis_layer2proxy_active")
	if err != nil {
		logs.Error("read redis::redis_layer2proxy_active failed, err:%v", err)
		return
	}

	//读取各类goroutine线程数量
	appConfig.ReadGoroutineNum, err = conf.Int("service::read_layer2proxy_goroutine_num")
	if err != nil {
		logs.Error("read service::read_layer2proxy_goroutine_num failed, err:%v", err)
		return
	}

	appConfig.WriteGoroutineNum, err = conf.Int("service::write_proxy2layer_goroutine_num")
	if err != nil {
		logs.Error("read service::write_proxy2layer_goroutine_num failed, err:%v", err)
		return
	}

	appConfig.HandleUserGoroutineNum, err = conf.Int("service::handle_user_goroutine_num")
	if err != nil {
		logs.Error("read service::handle_user_goroutine_num failed, err:%v", err)
		return
	}

	appConfig.Read2handleChanSize, err = conf.Int("service::read2handle_chan_size")
	if err != nil {
		logs.Error("read service::read2handle_chan_size failed, err:%v", err)
		return
	}

	appConfig.MaxRequestWaitTimeout, err = conf.Int("service::max_request_wait_timeout")
	if err != nil {
		logs.Error("read service::max_request_wait_timeout failed, err:%v", err)
		return
	}

	appConfig.Handle2WriteChanSize, err = conf.Int("service::handle2write_chan_size")
	if err != nil {
		logs.Error("read service::handle2write_chan_size failed, err:%v", err)
		return
	}

	appConfig.SendToWriteChanTimeout, err = conf.Int("service::send_to_write_chan_timeout")
	if err != nil {
		logs.Error("read service::send_to_write_chan_timeout failed, err:%v", err)
		return
	}

	appConfig.SendToHandleChanTimeout, err = conf.Int("service::send_to_handle_chan_timeout")
	if err != nil {
		logs.Error("read service::send_to_handle_chan_timeout failed, err:%v", err)
		return
	}

	//读取token秘钥
	appConfig.TokenPasswd = conf.String("service::seckill_token_passwd")
	if len(appConfig.TokenPasswd) == 0 {
		logs.Error("read service::seckill_token_passwd failed")
		err = fmt.Errorf("read service::seckill_token_passwd failed")
		return
	}

	//读取etcd相关的配置

	appConfig.EtcdConfig.EtcdAddr = conf.String("etcd::server_addr")
	if len(appConfig.TokenPasswd) == 0 {
		logs.Error("read service::seckill_token_passwd failed")
		err = fmt.Errorf("read service::seckill_token_passwd failed")
		return
	}

	etcdTimeout, err := conf.Int("etcd::etcd_timeout")
	if err != nil {
		err = fmt.Errorf("init config failed, read etcd_timeout error:%v", err)
		return
	}

	appConfig.EtcdConfig.Timeout = etcdTimeout
	appConfig.EtcdConfig.EtcdSecKeyPrefix = conf.String("etcd::etcd_sec_key_prefix")
	if len(appConfig.EtcdConfig.EtcdSecKeyPrefix) == 0 {
		err = fmt.Errorf("init config failed, read etcd_sec_key error:%v", err)
		return
	}

	productKey := conf.String("etcd::etcd_product_key")
	if len(productKey) == 0 {
		err = fmt.Errorf("init config failed, read etcd_product_key error:%v", err)
		return
	}

	if strings.HasSuffix(appConfig.EtcdConfig.EtcdSecKeyPrefix, "/") == false {
		appConfig.EtcdConfig.EtcdSecKeyPrefix = appConfig.EtcdConfig.EtcdSecKeyPrefix + "/"
	}

	appConfig.EtcdConfig.EtcdSecProductKey = fmt.Sprintf("%s%s", appConfig.EtcdConfig.EtcdSecKeyPrefix, productKey)
	return
}
