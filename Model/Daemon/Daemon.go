package Daemon

import (
	"../../util"
	"sync"
)

//任务基础接口
type task interface {
	getid() string //获取任务id
	start() error  //启动
	wait() error   //等待
	stop() error   //停止
}

//设置类型
type Config struct {
	TaskAccN  uint64
	TaskQSize uint64
}

var taskQ chan *task                  //任务队列，用以存储要执行的任务
var Qn = &count{0, new(sync.RWMutex)} //任务队列当前长度

func run1Task() {
	tsk := <-taskQ
	Qn.less() //队列中任务数量-1
	cancel, exists := cancelQ[tsk.getid()]
	if exists && cancel.get() > 0 { //如果取消就不运行
		cancel.less()
		return
	}
	if err := tsk.start(); err == nil { //否则就运行
		util.LogE(tsk.wait())
	} else { //运行出错则停止
		util.LogE(err)
		util.LogE(tsk.stop())
	}
}

//按照配置文件创建任务队列和执行任务的后台goroutine
func Init(Conf Config) {
	taskQ = make(chan *task, Conf.TaskQSize) //初始化任务队列
	for i := uint64(0); i < Conf.TaskAccN; i++ {
		go func() { //后台任务处理goroutine
			for {
				run1Task()
			}
		}()
	}
}

//将一个任务交给daemon运行
//
//入任务队列
func AddTask(tsk *task) {
	taskQ <- tsk //入队列
	Qn.more()    //队列中任务数量+1
}
