package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 选出一个新的分区(partition)
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success_channel 返回
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"                                   // 定义主题
	msg.Value = sarama.StringEncoder("this is a log info ") //日志内容

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config) // 连接卡夫卡
	if err != nil {
		fmt.Println("producer close error:", err)
		return
	}
	defer client.Close()

	pid, offset, err := client.SendMessage(msg) // 发送消息到卡夫卡
	if err != nil {
		fmt.Println("send msg error:", err)
		return
	}
	fmt.Println(pid, offset)

}
