package kafka

import (
	"sync"
	"sync/atomic"
	"time"
)

var (
	msgCount  int32
	msgLimit  int32 = 5000
	second    int64
	limitLock sync.Mutex
	lastFull  int64
)

const (
	SLEEP_MILLISEC = 50
)

func SetLimit(limit int32) bool {
	if limit > 0 && msgLimit != limit {
		msgLimit = limit
		return true
	}
	return false
}
func GetLimit() int32 {
	return msgLimit
}
func GetCount() int32 {
	return msgCount
}
func LimitDecr(delta int32) int32 {
	limitLock.Lock()
	defer limitLock.Unlock()
	if msgLimit-delta > 0 {
		msgLimit -= delta
	}
	return msgLimit
}
func LimitIncr(delta int32) int32 {
	limitLock.Lock()
	defer limitLock.Unlock()
	msgLimit += delta
	return msgLimit
}
func IsFull() bool {
	unixTime := time.Now().Unix()
	if lastFull+3 >= unixTime {
		return true
	}
	return false
}
func AddOne() {
	for {
		unixTime := time.Now().Unix()
		if unixTime <= second && msgCount >= msgLimit {
			time.Sleep(time.Duration(SLEEP_MILLISEC) * time.Millisecond)
			lastFull = unixTime
			continue
		}
		if unixTime > second {
			atomic.StoreInt32(&msgCount, 1)
			atomic.StoreInt64(&second, unixTime)
		} else {
			atomic.AddInt32(&msgCount, 1)
		}
		return
	}
}
