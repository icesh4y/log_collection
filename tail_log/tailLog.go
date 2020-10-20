package tail_log

import (
	"fmt"
	"github.com/hpcloud/tail"
)

// 收集日志。从日志文件读取日志内容

var tailObject *tail.Tail

func Init(fileName string)(err error)  {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件哪个位置开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,                                 //
	}
	tailObject, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail open file error,", err)
		return
	}
	return
}

func ReadChan() <-chan *tail.Line {
	return tailObject.Lines
}