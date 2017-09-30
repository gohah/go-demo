package main

import (
	"flag"
	"fmt"
	. "github.com/Shopify/sarama"
	"os"
)

func main() {
	offset := flag.Int64("offset", -1, "offset")
	topic := flag.String("topic", "b2c_applogs_acts_ctrl_logs_testing", "topic")
	partition := flag.Int("part", 0, "partition")
	flag.Parse()
	cfg := NewConfig()
	cfg.Consumer.Offsets.Initial = OffsetOldest
	client, err := NewConsumer([]string{"c3-b2c-systech-kafka01.bj:9092", "c3-b2c-systech-kafka02.bj:9092", "c3-b2c-systech-kafka03.bj:9092", "c3-b2c-systech-kafka04.bj:9092"}, cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	partiConsumer, err := client.ConsumePartition(*topic, int32(*partition), *offset)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < 10; i++ {
		msg := <-partiConsumer.Messages()
		fmt.Printf("%s\t%d\n", string(msg.Value), msg.Offset)
	}
}
