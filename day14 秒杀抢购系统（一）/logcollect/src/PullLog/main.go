package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	. "github.com/Shopify/sarama"
)

func GetDate(log string) string {
	arr := strings.Split(log, "]")
	return strings.Trim(arr[0], "[")
}
func GetHost(log string) (string, error) {
	arr := strings.Split(log, "]")
	if len(arr) < 3 {
		return "", fmt.Errorf("log:%s, error", log)
	}
	host := strings.TrimSpace(arr[2])
	host = strings.Trim(host, "[")
	return host, nil
}
func GetLog(topic string, partition int32, date string, start, end int64) {
	client, err := NewConsumer([]string{"localhost:9092"}, nil)
	partiConsumer, err := client.ConsumePartition(topic, partition, start-10000)
	if err != nil {
		fmt.Println(err, "get start")
		os.Exit(1)
	}

	out := make(map[string]*os.File)
	for {
		msg := <-partiConsumer.Messages()
		if msg.Offset > end+10000 {
			return
		}
		log := string(msg.Value)
		if strings.Contains(log, date) == false {
			continue
		}
		host, err := GetHost(log)
		if err != nil {
			f, _ := os.OpenFile(fmt.Sprintf("%s.error.log", topic), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			fmt.Fprintf(f, "log:%s, offset:%d, part:%d", log, msg.Offset, partition)
			continue
		}
		fp, ok := out[host]
		if !ok {
			f, _ := os.OpenFile(fmt.Sprintf("%s.%s.%s.log", topic, host, date), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			out[host] = f
		}
		fp = out[host]
		fmt.Fprintf(fp, "%s\n", log)
	}
}

func main() {
	topic := flag.String("topic", "nginx_log", "topic")
	partition := flag.Int("part", 0, "partition")
	date := flag.String("date", "2016-04-30", "date")
	start := flag.Int64("start", 0, "offset start")
	end := flag.Int64("end", 0, "offset end")
	flag.Parse()
	GetLog(*topic, int32(*partition), *date, *start, *end)
}
