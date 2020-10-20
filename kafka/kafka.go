package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// 讲日志内容写入卡夫卡

var client sarama.SyncProducer // 连接卡夫卡的生产者客户端

func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 选出一个新的分区(partition)
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success_channel 返回

	client, err = sarama.NewSyncProducer(addrs, config) // 连接卡夫卡
	if err != nil {
		fmt.Println("producer close error:", err)
		return
	}
	return
}

func SendToKafka(topic string, data string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic // 定义主题
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg) // 发送消息到卡夫卡
	if err != nil {
		fmt.Println("send msg error:", err)
		return
	}
	fmt.Println(pid, offset)

}
