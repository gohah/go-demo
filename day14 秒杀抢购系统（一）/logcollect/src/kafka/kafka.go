package kafka

import (
	"github.com/Saerdna/golib/common"
	. "github.com/Shopify/sarama"
	"time"
	"fmt"
	"os"
)

var (
	producer     AsyncProducer
	globalConfig *KafkaConfig
	servers      []string
)

type KafkaConfig struct {
	MaxMessageBytes int              //single msg byte size , default 1MB
	CompressionCode CompressionCodec //compressioncode, default none compress
	Frequency       time.Duration    //msg buffer time
	MaxMessages     int              //max buffer msg nums
	ClientID        string
}

func ConfigInit() {
	globalConfig = &KafkaConfig{
		MaxMessageBytes: 1024 * 1024,
		CompressionCode: CompressionNone,
		Frequency:       time.Duration(2) * time.Second,
		MaxMessages:     4000,
		ClientID:        common.GetLocalIP(),
	}
}
func SetMaxMsgByte(MaxBytes int) bool {
	if MaxBytes > 0 && globalConfig.MaxMessageBytes != MaxBytes {
		globalConfig.MaxMessageBytes = MaxBytes
		return true
	}
	return false
}
func SetCompressType(Type int) bool {
	CompressType := CompressionCodec(Type)
	if CompressType == globalConfig.CompressionCode {
		return false
	}
	switch CompressType {
	case CompressionNone:
		globalConfig.CompressionCode = CompressionNone
	case CompressionGZIP:
		globalConfig.CompressionCode = CompressionGZIP
	case CompressionSnappy:
		globalConfig.CompressionCode = CompressionSnappy
	default:
		globalConfig.CompressionCode = CompressionNone
	}
	return true
}
func SetFreqTime(Second int) bool {
	if globalConfig.Frequency != time.Duration(Second)*time.Second {
		globalConfig.Frequency = time.Duration(Second) * time.Second
		return true
	}
	return false
}
func SetMaxMsgs(Num int) bool {
	if globalConfig.MaxMessages != Num {
		globalConfig.MaxMessages = Num
		return true
	}
	return false
}
func GetServers() []string {
	return servers
}
func Connect(addrs []string) error {
	if globalConfig == nil {
		ConfigInit()
	}
	config := NewConfig()
	config.Producer.MaxMessageBytes = globalConfig.MaxMessageBytes
	config.Producer.Compression = globalConfig.CompressionCode
	config.Producer.Flush.Frequency = globalConfig.Frequency
	config.Producer.Flush.MaxMessages = globalConfig.MaxMessages
	config.ClientID = globalConfig.ClientID
	servers = addrs
	client, err := NewAsyncProducer(addrs, config)
	if err == nil {
		producer = client
	}
    go handleErrors()
	return err
}
func Close() error {
	return producer.Close()
}
func SendMessageByChan(topic string, msgCh chan string) error {
	msgObj := new(ProducerMessage)
	msgObj.Topic = topic
	msgObj.Key = StringEncoder(globalConfig.ClientID)
	for msg := range msgCh {
		msgObj.Value = StringEncoder(msg)
	    producer.Input() <- msgObj
	}
	return nil
}

func SendMessage(topic, msg string) error {
	AddOne()
	select {
	case producer.Input() <- &ProducerMessage{Topic: topic, Value: StringEncoder(msg)}:
		return nil
	case err := <-producer.Errors():
		return err
	}
}

func handleErrors() {
	defer recover()
	for item := range producer.Errors() {
		fmt.Fprintf(os.Stderr, "kafka send error:%v", item)
	}
}
