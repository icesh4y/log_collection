package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"log_collection/kafka"
	"log_collection/tail_log"
	"time"
)

func run(topic string) {
	for {
		select {
		case lines := <-tail_log.ReadChan():
			kafka.SendToKafka(topic, lines.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 1.初始化卡夫卡连接
	// 2.打开日志文件 准备读入内容

	cfg, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Println("ini load error:", err)
		return
	}

	if err := kafka.Init([]string{cfg.Section("kafka").Key("address").String()}); err != nil {
		fmt.Println("kafka init error:", err)
		return
	}

	if err := tail_log.Init(cfg.Section("taillog").Key("path").String()); err != nil {
		fmt.Println("tail open file error:", err)
		return
	}

	run(cfg.Section("kafka").Key("topic").String())

}
