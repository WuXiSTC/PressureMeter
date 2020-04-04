package Daemon

import (
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"sync"
	"time"
)

//任务基础接口
type TaskInterface interface {
	GetID() string                                           //获取任务ID
	Start(shutdownPort uint16, duration time.Duration) error //启动
	Wait()                                                   //等待
	Stop() error                                             //停止
}

type task struct {
	TaskInterface
	duration time.Duration
}

var taskQ chan *task                //任务队列，用以存储要执行的任务的地址
var Qn = &count{0, new(sync.Mutex)} //任务队列当前长度

//运行一个任务
func run1task(i uint16) {
	tsk := <-taskQ
	if tsk == nil {
		return
	}
	util.Log(fmt.Sprintf("Daemon %d: get task %s", i, tsk.GetID()))
	cancelQMu.RLock()
	Qn.less() //队列中任务数量-1
	cancel, exists := cancelQ[tsk.GetID()]
	cancelQMu.RUnlock()
	if exists && cancel.less() >= 0 { //如果已取消就不运行
		return
	}
	if err := tsk.Start(conf.BasePort-i, tsk.duration); err == nil { //否则就运行
		util.Log(fmt.Sprintf("Daemon %d: started task %s", i, tsk.GetID()))
		tsk.Wait()
		util.Log(fmt.Sprintf("Daemon %d: stopped task %s", i, tsk.GetID()))
	} else { //运行出错则停止
		util.LogE(err)
		err = tsk.Stop()
		util.LogE(err)
	}
}

var toStop = false
var stopped = make(chan uint16)

//将一个任务交给daemon运行
//
//入任务队列
func AddTask(tsk TaskInterface, duration time.Duration) {
	taskQ <- &task{
		TaskInterface: tsk,
		duration:      duration,
	} //入队列
	Qn.more() //队列中任务数量+1
}

//停止Daemon运行
func Stop() {
	toStop = true
	close(taskQ)
	for i := uint16(0); i < conf.TaskAccN; i++ {
		goi := <-stopped
		util.Log(fmt.Sprintf("Daemon %d stopped", goi))
	}
}

func ExpectDuration() time.Duration {
	d := time.Duration(0)
	return d / time.Duration(conf.TaskAccN)
}
