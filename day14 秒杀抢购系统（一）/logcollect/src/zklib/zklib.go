package zklib

import (
	"encoding/json"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"path"
	"strings"
	"time"
)

const (
	PATH_PREFIX  = "/LogCollect/Nodes/"
	KAFKA_PREFIX = "/LogCollect/Kafka/"
)

type TopicInfo struct {
	TopicName  string
	FilePath   string
	PrefixType int
	SplitType  int
}

type NodeConfig struct {
	Cluster      string
	TicketTime   int64
	MaxBytes     int
	CompressType int
	Frequency    int
	MaxMsgs      int
	MsgLimit     int32
	LimitType    int
}

func DealSession(session <-chan zk.Event) {
	for e := range session {
		if e.State == zk.StateDisconnected {
			return
		}
	}
}

var (
	zkCli *zk.Conn
)

func Connect(servers []string, timeout int, logger zk.Logger) error {
	client, session, err := zk.Connect(servers, time.Duration(timeout)*time.Second)
	if err != nil {
		return fmt.Errorf("servers:%+v err:%s", servers, err)
	}
	go DealSession(session)
	if logger != nil {
		client.SetLogger(logger)
	}
	zkCli = client
	return nil
}
func GetKafkaConfig(cluster string) ([]string, error) {
	data, err := GetNodeData(path.Join(KAFKA_PREFIX, cluster))
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), ","), nil
}
func GetNodeConfig(addr string) (*NodeConfig, error) {
	data, err := GetNodeData(path.Join(PATH_PREFIX, addr))
	if err != nil {
		return nil, err
	}
	var obj NodeConfig
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, fmt.Errorf("node:%s err:%s", path.Join(PATH_PREFIX, addr), err)
	}
	return &obj, nil
}
func GetTopicConfig(addr, topicname string) (*TopicInfo, error) {
	data, err := GetNodeData(path.Join(PATH_PREFIX, addr, topicname))
	if err != nil {
		return nil, err
	}
	var obj TopicInfo
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, fmt.Errorf("node:%s, err:%s", path.Join(PATH_PREFIX, addr, topicname), err)
	}
	return &obj, nil
}
func GetNodeData(nodepath string) ([]byte, error) {
	data, _, err := zkCli.Get(nodepath)
	if err != nil {
		return nil, fmt.Errorf("node:%s, err:%s", nodepath, err)
	}
	return data, nil
}
func GetNodeChildren(addr string) ([]string, error) {
	children, _, err := zkCli.Children(path.Join(PATH_PREFIX, addr))
	if err != nil {
		return children, fmt.Errorf("path:%s, err:%s", path.Join(PATH_PREFIX, addr), err)
	}
	return children, nil
}
func WatchNodeEvent(addr string) (<-chan zk.Event, error) {
	_, _, eventCh, err := zkCli.GetW(path.Join(PATH_PREFIX, addr))
	if err != nil {
		return nil, fmt.Errorf("path:%s, err:%s", path.Join(PATH_PREFIX, addr), err)
	}
	return eventCh, nil
}
func WatchTopicEvent(addr, topicname string) (<-chan zk.Event, error) {
	_, _, eventCh, err := zkCli.GetW(path.Join(PATH_PREFIX, addr, topicname))
	if err != nil {
		return nil, fmt.Errorf("path:%s, err:%s", path.Join(PATH_PREFIX, addr, topicname), err)
	}
	return eventCh, err

}
func WatchChildren(addr string) (<-chan zk.Event, error) {
	_, _, eventCh, err := zkCli.ChildrenW(path.Join(PATH_PREFIX, addr))
	if err != nil {
		return nil, fmt.Errorf("path:%s, err:%s", path.Join(PATH_PREFIX, addr), err)
	}
	return eventCh, err
}
