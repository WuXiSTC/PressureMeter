package Daemon

import (
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"github.com/yindaheng98/go-utility/QueueSet"
	"strconv"
	"time"
)

//设置类型
type Config struct {
	TaskAccN  uint16        `yaml:"TaskAccN" usage:"以同时进行的任务数量"`
	TaskQSize uint16        `yaml:"TaskQSize" usage:"任务队列缓冲区大小"`
	RestTime  time.Duration `yaml:"RestTime" usage:"线程在前一个任务任务结束到后一个任务开始之间的休息时间"`
}

var conf Config //配置信息

func DefaultConfig() Config {
	return Config{
		TaskAccN:  4,
		TaskQSize: 100,
		RestTime:  5e9,
	}
}

var toStop = false
var stopped = make(chan uint16)

//按照配置文件创建任务队列和执行任务的后台goroutine
func Init(c Config) {
	conf = c
	util.Log(fmt.Sprintf("%d tasks can running simultaneously at most", conf.TaskAccN))
	util.Log(fmt.Sprintf("Task buffer size: %d", conf.TaskQSize))
	queue = QueueSet.New(uint64(conf.TaskQSize))
	runnings = make([]TaskInterface, conf.TaskAccN)
	for i := uint16(0); i < conf.TaskAccN; i++ {
		go func(goi uint16) { //后台任务处理goroutine
			for !toStop { //如果检测到要停止了就停止
				run1task(goi)
				time.Sleep(conf.RestTime)
			}
			stopped <- goi
		}(i)
		util.Log("Daemon " + strconv.Itoa(int(i)) + " started")
	}
}

//停止Daemon运行
func Stop() {
	toStop = true
	for i := uint16(0); i < conf.TaskAccN; i++ {
		goi := <-stopped
		util.Log(fmt.Sprintf("Daemon %d stopped", goi))
	}
}
