package service

import (
	"sync"
	"time"

	etcd_client "github.com/coreos/etcd/clientv3"
	"github.com/garyburd/redigo/redis"
)

var (
	secLayerContext = &SecLayerContext{}
)

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
	//限速控制
	secLimit *SecLimit
}

type RedisConf struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
	RedisQueueName   string
}

type EtcdConf struct {
	EtcdAddr          string
	Timeout           int
	EtcdSecKeyPrefix  string
	EtcdSecProductKey string
}

type SecLayerConf struct {
	Proxy2LayerRedis RedisConf
	Layer2ProxyRedis RedisConf
	EtcdConfig       EtcdConf
	LogPath          string
	LogLevel         string

	WriteGoroutineNum      int
	ReadGoroutineNum       int
	HandleUserGoroutineNum int
	Read2handleChanSize    int
	Handle2WriteChanSize   int
	MaxRequestWaitTimeout  int

	SendToWriteChanTimeout  int
	SendToHandleChanTimeout int

	SecProductInfoMap map[int]*SecProductInfoConf
	TokenPasswd       string
}

type SecLayerContext struct {
	proxy2LayerRedisPool *redis.Pool
	layer2ProxyRedisPool *redis.Pool
	etcdClient           *etcd_client.Client
	RWSecProductLock     sync.RWMutex

	secLayerConf     *SecLayerConf
	waitGroup        sync.WaitGroup
	Read2HandleChan  chan *SecRequest
	Handle2WriteChan chan *SecResponse

	HistoryMap     map[int]*UserBuyHistory
	HistoryMapLock sync.Mutex

	//商品的计数
	productCountMgr *ProductCountMgr
}

type SecRequest struct {
	ProductId     int
	Source        string
	AuthCode      string
	SecTime       string
	Nance         string
	UserId        int
	UserAuthSign  string
	AccessTime    time.Time
	ClientAddr    string
	ClientRefence string
	//CloseNotify   <-chan bool

	//ResultChan chan *SecResult
}

type SecResponse struct {
	ProductId int
	UserId    int
	Token     string
	TokenTime int64
	Code      int
}
