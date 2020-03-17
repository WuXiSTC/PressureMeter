package Daemon

import (
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"strconv"
)

//设置类型
type Config struct {
	TaskAccN  uint64 //以同时进行的任务数量
	TaskQSize uint64 //任务队列缓冲区大小
}

var conf Config //配置信息

func DefaultConfig() Config {
	return Config{
		TaskAccN:  4,
		TaskQSize: 100,
	}
}

//按照配置文件创建任务队列和执行任务的后台goroutine
func Init(c Config) {
	conf = c
	util.Log(fmt.Sprintf("%d tasks can running simultaneously at most", conf.TaskAccN))
	util.Log(fmt.Sprintf("Task buffer size: %d", conf.TaskQSize))
	taskQ = make(chan *TaskInterface, conf.TaskQSize) //初始化任务队列
	for i := uint64(0); i < conf.TaskAccN; i++ {
		go func(goi uint64) { //后台任务处理goroutine
			for !toStop { //如果检测到要停止了就停止
				run1task(goi)
			}
			stopped <- goi
		}(i)
		util.Log("Daemon " + strconv.Itoa(int(i)) + " started")
	}
}
