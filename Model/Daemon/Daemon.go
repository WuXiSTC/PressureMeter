package Daemon

import (
	"../../util"
	"sync"
)

//任务基础接口
type TaskInterface interface {
	GetID() string //获取任务ID
	Start() error  //启动
	Wait() error   //等待
}

//设置类型
type Config struct {
	TaskAccN  uint64 `yaml:"TaskAccN"`
	TaskQSize uint64 `yaml:"TaskQSize"`
}

var taskQ chan *TaskInterface         //任务队列，用以存储要执行的任务的地址
var Qn = &count{0, new(sync.RWMutex)} //任务队列当前长度

//运行一个任务
func run1task() {
	tsk := <-taskQ
	Qn.less() //队列中任务数量-1
	cancel, exists := cancelQ[(*tsk).GetID()]
	if exists && cancel.get() > 0 { //如果取消就不运行
		cancel.less()
		return
	}
	if err := (*tsk).Start(); err == nil { //否则就运行
		util.LogE((*tsk).Wait())
	} else { //运行出错则停止
		util.LogE(err)
	}
}

var toStop = false

//按照配置文件创建任务队列和执行任务的后台goroutine
func Init(Conf Config) {
	taskQ = make(chan *TaskInterface, Conf.TaskQSize) //初始化任务队列
	for i := uint64(0); i < Conf.TaskAccN; i++ {
		go func() { //后台任务处理goroutine
			for !toStop { //如果检测到要停止了就停止
				run1task()
			}
		}()
	}
}

//将一个任务交给daemon运行
//
//入任务队列
func AddTask(tsk *TaskInterface) {
	taskQ <- tsk //入队列
	Qn.more()    //队列中任务数量+1
}

//停止Daemon运行
func Stop() {
	toStop = true
}
