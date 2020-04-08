package Daemon

import (
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"github.com/yindaheng98/go-utility/QueueSet"
	"sync"
	"time"
)

var queue *QueueSet.QueueSet

//将一个任务交给daemon运行
//
//入任务队列
func Queue(t TaskInterface) {
	runningsMu.RLock()
	defer runningsMu.RUnlock()
	queue.Push(t)
	conf.UpdateStateCallback(getInfo())
}

//取消一个任务的运行
func Cancel(id string) {
	runningsMu.RLock()
	defer runningsMu.RUnlock()
	queue.Cancel(id)
	for i, task := range runnings {
		if task != nil && task.GetID() == id {
			task.Stop(uint16(i))
		}
	}
	conf.UpdateStateCallback(getInfo())
}

var runnings []TaskInterface
var startTimes []time.Time
var runningsMu = new(sync.RWMutex)

//运行一个任务
func run1task(i uint16) {
	task := queue.Pop().(TaskInterface) //出队列
	if task == nil {
		return
	}

	runningsMu.Lock()
	runnings[i] = task
	startTimes[i] = time.Now()
	conf.UpdateStateCallback(getInfo())
	util.Log(fmt.Sprintf("Daemon %d: get task %s", i, task.GetID()))
	task.Start(i) //启动
	util.Log(fmt.Sprintf("Daemon %d: started task %s", i, task.GetID()))
	runningsMu.Unlock()

	task.Wait() //等待完成

	runningsMu.Lock()
	util.Log(fmt.Sprintf("Daemon %d: stopped task %s", i, task.GetID()))
	runnings[i] = nil
	conf.UpdateStateCallback(getInfo())
	runningsMu.Unlock()
}

func getInfo() (rs []TaskInterface, ts []time.Time, qs []TaskInterface) {
	rs, ts = make([]TaskInterface, len(runnings)), make([]time.Time, len(startTimes))
	for i, t := range runnings {
		rs[i], ts[i] = t, startTimes[i]
	}

	es := queue.GetQueueElements()
	qs = make([]TaskInterface, len(es))
	for i, t := range es {
		qs[i] = t.(TaskInterface)
	}
	return
}
