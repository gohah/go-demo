package main

import (
	"flag"
	"fmt"
	. "github.com/Shopify/sarama"
	"os"
	"time"
)

func TopicMonitor(addrs []string, topic string) {
	client, _ := NewClient(addrs, nil)
	partids, _ := client.Partitions(topic)
	poffset := make(map[int32]int64)
	noffset := make(map[int32]int64)
	var total int64
	for {
		total = 0
		for _, partid := range partids {
			offset, _ := client.GetOffset(topic, int32(partid), OffsetNewest)
			noffset[partid] = offset
		}
		for _, partid := range partids {
			if _, ok := poffset[partid]; ok {
				total += noffset[partid] - poffset[partid]
			}
			poffset[partid] = noffset[partid]
		}
		time.Sleep(1 * time.Second)
		if total == 0 {
			continue
		}
		fmt.Printf("topic:%s, %d msg/s\n", topic, total)
	}
}
func main() {
	host := flag.String("host", "127.0.0.1:9092", "host")
	flag.Parse()
	client, err := NewClient([]string{*host}, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(client.Topics())
	topics, _ := client.Topics()
	for _, topic := range topics {
		go TopicMonitor([]string{*host}, topic)
	}
	time.Sleep(100 * time.Second)
}
