package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件哪个位置开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,                                 //
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail open file error,", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Println("tails.lines read nil ,filename :", fileName)
			time.Sleep(time.Second)
		}
		fmt.Println(line.Text)
	}
}
