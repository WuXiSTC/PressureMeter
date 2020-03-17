package Daemon

import (
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"strconv"
	"sync"
)

//任务基础接口
type TaskInterface interface {
	GetID() string //获取任务ID
	Start() error  //启动
	Wait() error   //等待
	Stop() error   //停止
}

var taskQ chan *TaskInterface         //任务队列，用以存储要执行的任务的地址
var Qn = &count{0, new(sync.RWMutex)} //任务队列当前长度

//运行一个任务
func run1task(i uint64) {
	tsk := <-taskQ
	if tsk == nil {
		return
	}
	util.Log("Daemon " + strconv.Itoa(int(i)) + ": get " + (*tsk).GetID())
	Qn.less() //队列中任务数量-1
	cancel, exists := cancelQ[(*tsk).GetID()]
	if exists && cancel.get() > 0 { //如果取消就不运行
		cancel.less()
		return
	}
	if err := (*tsk).Start(); err == nil { //否则就运行
		util.Log(fmt.Sprintf("Daemon %d: started task %s", i, (*tsk).GetID()))
		util.LogE((*tsk).Wait())
		util.Log(fmt.Sprintf("Daemon %d: stopped task %s", i, (*tsk).GetID()))
	} else { //运行出错则停止
		util.LogE(err)
		err = (*tsk).Stop()
		util.LogE(err)
	}
}

var toStop = false
var stopped = make(chan uint64)

//将一个任务交给daemon运行
//
//入任务队列
func AddTask(tsk TaskInterface) {
	taskQ <- &tsk //入队列
	Qn.more()     //队列中任务数量+1
}

//停止Daemon运行
func Stop() {
	toStop = true
	close(taskQ)
	for i := uint64(0); i < conf.TaskAccN; i++ {
		goi := <-stopped
		util.Log("Daemon " + strconv.Itoa(int(goi)) + " stopped")
	}
}
