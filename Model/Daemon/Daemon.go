package Daemon

import (
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"github.com/yindaheng98/go-utility/QueueSet"
	"sync"
)

var queue *QueueSet.QueueSet

//将一个任务交给daemon运行
//
//入任务队列
func Queue(t TaskInterface) {
	queue.Push(t)
}

//取消一个任务的运行
func Cancel(id string) {
	queue.Cancel(id)
	runningsMu.RLock()
	defer runningsMu.RUnlock()
	for i, task := range runnings {
		if task != nil && task.GetID() == id {
			task.Stop(uint16(i))
		}
	}
}

var runnings []TaskInterface
var runningsMu = new(sync.RWMutex)

//运行一个任务
func run1task(i uint16) {
	task := queue.Pop().(TaskInterface) //出队列
	if task == nil {
		return
	}

	runningsMu.Lock()
	runnings[i] = task
	util.Log(fmt.Sprintf("Daemon %d: get task %s", i, task.GetID()))
	task.Start(i) //启动
	util.Log(fmt.Sprintf("Daemon %d: started task %s", i, task.GetID()))
	runningsMu.Unlock()

	task.Wait() //等待完成

	runningsMu.Lock()
	util.Log(fmt.Sprintf("Daemon %d: stopped task %s", i, task.GetID()))
	runnings[i] = nil
	runningsMu.Unlock()
}
