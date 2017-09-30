package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/Shopify/sarama"
)

func GetDate(log string) string {
	arr := strings.Split(log, "]")
	return strings.Trim(arr[0], "[")
}
func GetHost(log string) string {
	arr := strings.Split(log, "]")
	host := strings.Trim(arr[0], "[")
	host = strings.TrimSpace(host)
	return host
}
func GetOffset(topic string, partition int32, date string) (start int64, end int64) {
	client, err := NewConsumer([]string{"localhost:9092"}, nil)
	partiConsumer, err := client.ConsumePartition(topic, partition, OffsetOldest)
	if err != nil {
		fmt.Println(err, "get low")
		os.Exit(1)
	}
	var low int64
	var high int64
	select {
	case msg := <-partiConsumer.Messages():
		low = msg.Offset
	case <-time.After(2 * time.Second):
		fmt.Println("get low failed")
		return -1, -1
	}
	partiConsumer.Close()
	partiConsumer, err = client.ConsumePartition(topic, partition, OffsetNewest)
	if err != nil {
		fmt.Println(err, "get high")
		os.Exit(1)
	}
	select {
	case msg := <-partiConsumer.Messages():
		high = msg.Offset
	case <-time.After(100 * time.Second):
		fmt.Println("get high failed")
		return -1, -1
	}
	partiConsumer.Close()
	var mid int64
	var l int64
	var r int64
	l = low
	r = high
	for {
		if low > high {
			break
		}
		mid = (low + high) / 2
		partiConsumer, err = client.ConsumePartition(topic, partition, mid)
		if err != nil {
			fmt.Println(err, "get start")
			os.Exit(1)
		}
		msg := <-partiConsumer.Messages()
		partiConsumer.Close()
		logDate := GetDate(string(msg.Value))
		fmt.Printf("%s--%s--%d--%d\n", logDate, date, mid, msg.Offset)
		switch {
		case strings.Contains(logDate, date):
			start = mid
			high = mid - 1
		case logDate < date:
			low = mid + 1
		case logDate > date:
			high = mid - 1
		}
	}
	low = l
	high = r
	for {
		if low > high {
			break
		}
		mid = (low + high) / 2
		partiConsumer, err = client.ConsumePartition(topic, partition, mid)
		if err != nil {
			fmt.Println(err, "get end")
			os.Exit(1)
		}
		msg := <-partiConsumer.Messages()
		partiConsumer.Close()
		logDate := GetDate(string(msg.Value))
		fmt.Printf("%s--%s--%d--%d\n", logDate, date, mid, msg.Offset)
		switch {
		case strings.Contains(logDate, date):
			end = mid
			low = mid + 1
		case logDate < date:
			low = mid + 1
		case logDate > date:
			high = mid - 1
		}
	}
	return start, end
}

func main() {
	topic := flag.String("topic", "b2c_applogs_acts_ctrl_logs_testing", "topic")
	partition := flag.Int("part", 0, "partition")
	date := flag.String("date", "2016-05-11 09:50:00", "date")
	flag.Parse()
	fmt.Println(GetOffset(*topic, int32(*partition), *date))
}
