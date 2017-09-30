package main

import (
	"fmt"
	"math/rand"
	"oldboy/xstatsd"
	"runtime/debug"
	"strings"
	"time"
	"util"

	"oldboy/logcollect/src/kafka"

	"oldboy/logcollect/src/zklib"

	"oldboy/logcollect/src/tailf"

	"oldboy/golib/cmp"

	"github.com/samuel/go-zookeeper/zk"
)

var (
	ScheCh     map[string]chan Command
	TopicList  []string
	ticketTime time.Duration
	recvCmd    chan Command
)

const (
	TOPIC_START = iota
	TOPIC_STOP
	NODE_STOP
	NODE_UPDATE
	NODE_SHUTDOWN
)
const (
	WATCH_TOPIC = iota
	WATCH_NODE
	WATCH_CHILD
)

type Command struct {
	Cmd        int
	Topic      string
	Path       string
	PrefixType int
	SplitType  int
}

func NodeShutDown() {
	recvCmd <- Command{Cmd: NODE_SHUTDOWN}
}

func ReloadConfig(ip string) error {
	nodeCfg, err := zklib.GetNodeConfig(ip)
	for err != nil {
		<-time.After(time.Second * time.Duration(60+rand.Int()%10))
		nodeCfg, err = zklib.GetNodeConfig(ip)
	}
	if nodeCfg.TicketTime != 0 {
		ticketTime = time.Duration(nodeCfg.TicketTime)
	}
	var valid bool = false
	valid = kafka.SetMaxMsgByte(nodeCfg.MaxBytes) || valid
	valid = kafka.SetCompressType(nodeCfg.CompressType) || valid
	valid = kafka.SetFreqTime(nodeCfg.Frequency) || valid
	valid = kafka.SetMaxMsgs(nodeCfg.MaxMsgs) || valid
	valid = kafka.SetLimit(nodeCfg.MsgLimit) || valid
	limitType = nodeCfg.LimitType
	if ticketTime == 0 {
		ticketTime = time.Duration(10)
	}
	servers, err := zklib.GetKafkaConfig(nodeCfg.Cluster)
	if err == nil && servers != nil && len(servers) > 0 && valid && cmp.DiffSliceString(servers, kafka.GetServers()) {
		if len(servers) == 1 && (servers[0] == "\"\"" || servers[0] == "") {
			return nil
		}
		kafka.Connect(servers)
		return nil
	}
	return err
}
func TaskMain(cmd chan Command) {
	defer func() {
		if e := recover(); e != nil {
			taskLog.Warn("watch data panic, err:%s, stack:%s", e, debug.Stack())
		}
	}()
	item, ok := <-cmd
	if !ok {
		return
	}
	topic := item.Topic
	tailObj, err := tailf.NewTail(item.Path, item.SplitType, item.PrefixType, topic)
	standardIp := strings.Replace(util.GetLocalIP(), ".", "_", -1)
	metricMachineName := fmt.Sprintf("msg.machine.%v.%v.send", topic, standardIp)
	standardTopic := strings.Replace(topic, ".", "_", -1)
	metricTopicName := fmt.Sprintf("msg.topic.%v.send", standardTopic)
	ticket := time.NewTicker(ticketTime * time.Second)
	defer func() {
		ticket.Stop()
	}()
	monitorTicket := time.NewTicker(10 * time.Second)
	defer func() {
		monitorTicket.Stop()
	}()
	peroidCount := int64(0)
	secondCount := int64(0)
	for {
		select {
		case item, ok := <-cmd:
			if !ok {
				taskLog.Warn("cmd close, topic:%s close", topic)
				return
			}
			if item.Cmd == TOPIC_STOP {
				if err := tailObj.Stop(); err != nil {
					taskLog.Warn("topic:%s tail stop error:%s", topic, err)
				}
				return
			}
		case log, ok := <-tailObj.GetOneLine():
			if !ok {
				taskLog.Warn("topic:%s tail get line error:%s", topic, err)
				continue
			}
			if err := kafka.SendMessage(topic, log.Text); err != nil {
				taskLog.Warn("topic:%s SendMessage error:%s", topic, err)
			}
			peroidCount++
			secondCount++
		case <-ticket.C:
			xstatsd.Count(metricTopicName, secondCount)
			xstatsd.Count(metricMachineName, secondCount)
			secondCount = 0
			if err := tailObj.FlushOffset(); err != nil {
				taskLog.Warn("topic:%s flush offset error:%s", topic, err)
			}
		case <-monitorTicket.C:
			taskLog.Monitor(LogMap{"topic": topic, "status": "alive", "count": peroidCount})
			peroidCount = 0
		}
	}
}

func watchData(recv chan<- Command, addr, topic string, watchType int) {
	defer func() {
		if e := recover(); e != nil {
			taskLog.Warn("watch data panic, err:%s, stack:%s", e, debug.Stack())
		}
	}()
	var eventCh <-chan zk.Event
	var err error
	for {
		switch watchType {
		case WATCH_TOPIC:
			eventCh, err = zklib.WatchTopicEvent(addr, topic)
			if err != nil {
				taskLog.Warn("zk watch topic:%s err:%s", topic, err)
				continue
			}
		case WATCH_NODE:
			eventCh, err = zklib.WatchNodeEvent(addr)
			if err != nil {
				taskLog.Warn("zk watch topic:%s err:%s", topic, err)
				continue
			}
		case WATCH_CHILD:
			eventCh, err = zklib.WatchChildren(addr)
			if err != nil {
				taskLog.Warn("zk watch topic:%s err:%s", topic, err)
				continue
			}
		}
		event := <-eventCh
		if event.Err != nil {
			taskLog.Warn("watch event error:%s", event.Err)
			continue
		}
		switch event.Type {
		case zk.EventNodeDeleted:
			if watchType == WATCH_NODE {
				recv <- Command{Cmd: NODE_STOP}
			} else if watchType == WATCH_TOPIC {
				recv <- Command{Cmd: TOPIC_STOP, Topic: topic}
			}
			return
		case zk.EventNodeDataChanged:
			if watchType == WATCH_NODE {
				recv <- Command{Cmd: NODE_UPDATE}
			} else if watchType == WATCH_TOPIC {
				recv <- Command{Cmd: TOPIC_STOP, Topic: topic}
				recv <- Command{Cmd: TOPIC_START, Topic: topic}
				return
			}
		case zk.EventNodeChildrenChanged:
			recv <- Command{Cmd: NODE_UPDATE}
		}
	}
}
func TopicStart(addr, topic string) {
	topicCfg, err := zklib.GetTopicConfig(addr, topic)
	if err != nil {
		taskLog.Warn("Get Topic:%s config err:%s", topic, err)
		return
	}
	taskLog.Monitor(LogMap{"command": "start", "topic": topic, "path": topicCfg.FilePath, "prefixType": topicCfg.PrefixType, "splitType": topicCfg.SplitType})
	//taskLog.Notice("topic:%s start, path:%s, PrefixType:%d, SplitType:%d", topic, topicCfg.FilePath, topicCfg.PrefixType, topicCfg.SplitType)
	if _, ok := ScheCh[topic]; !ok {
		ScheCh[topic] = make(chan Command)
	}
	go TaskMain(ScheCh[topic])
	ScheCh[topic] <- Command{Cmd: TOPIC_START,
		Topic:      topic,
		Path:       topicCfg.FilePath,
		PrefixType: topicCfg.PrefixType,
		SplitType:  topicCfg.SplitType,
	}
}
func TopicStop(item Command) {
	taskLog.Monitor(LogMap{"command": "stop", "topic": item.Topic})
	if _, ok := ScheCh[item.Topic]; !ok {
		return
	}
	ScheCh[item.Topic] <- item
	close(ScheCh[item.Topic])
	delete(ScheCh, item.Topic)
}
func watchLimit() {
	for {
		time.Sleep(10 * time.Second)
		taskLog.Monitor(LogMap{"msgLimit": kafka.GetLimit(), "msgCount": kafka.GetCount(), "totalCpu": totalCpuUsage, "procCpu": procCpuUsage, "limitType": limitType})
	}
}
func watchErrLog() {
	for {
		select {
		case log, ok := <-tailf.ErrMsgCh:
			if !ok {
				time.Sleep(time.Second)
				continue
			}
			taskLog.Warn(log.Error())
		}
	}
}
func Gao() {
	localIP := util.GetLocalIP()
	taskLog.Monitor(LogMap{"startup": localIP})
	var err error
	if err = ReloadConfig(localIP); err != nil {
		taskLog.Warn("node config init failed, %s", err)
	}

	TopicList, err = zklib.GetNodeChildren(localIP)
	if err != nil {
		taskLog.Warn("get topic init failed, %s", err)
	}
	ScheCh = make(map[string]chan Command)
	recvCmd = make(chan Command)
	go RunSumLog()
	go watchLimit()
	go watchErrLog()
	go watchData(recvCmd, localIP, "", WATCH_NODE)
	go watchData(recvCmd, localIP, "", WATCH_CHILD)
	for _, topic := range TopicList {
		TopicStart(localIP, topic)
		go watchData(recvCmd, localIP, topic, WATCH_TOPIC)
	}
	for {
		item, ok := <-recvCmd
		if !ok {
			return
		}
		taskLog.Notice("recv command cmd:%d, topic:%s, path:%s", item.Cmd, item.Topic, item.Path)
		switch item.Cmd {
		case TOPIC_STOP:
			TopicStop(item)
		case TOPIC_START:
			TopicStart(localIP, item.Topic)
			go watchData(recvCmd, localIP, item.Topic, WATCH_TOPIC)
		case NODE_UPDATE:
			if err = ReloadConfig(localIP); err != nil {
				taskLog.Warn("node config reload failed, %s", err)
			}
			newTopicList, err := zklib.GetNodeChildren(localIP)
			if err != nil {
				taskLog.Warn("Get Node Child err:%s", err)
				continue
			}
			more, less := cmp.Compare(TopicList, newTopicList)
			for _, topic := range more {
				TopicStop(Command{Cmd: TOPIC_STOP, Topic: topic})
			}
			for _, topic := range less {
				TopicStart(localIP, topic)
				go watchData(recvCmd, localIP, topic, WATCH_TOPIC)
			}
			TopicList = newTopicList
		case NODE_STOP:
			for _, topic := range TopicList {
				TopicStop(Command{Cmd: TOPIC_STOP, Topic: topic})
			}
		case NODE_SHUTDOWN:
			for _, topic := range TopicList {
				TopicStop(Command{Cmd: TOPIC_STOP, Topic: topic})
			}
			taskLog.Fatal("logcollect stop")
			kafka.Close()
			return
		}
	}
}
