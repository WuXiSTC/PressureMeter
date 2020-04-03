package TaskList

import (
	"errors"
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	"time"
)

//将一个任务加进任务队列
//
//不会返回错误，返回的bool表示任务是否存在
func (tasklist *taskList) Start(id string, duration time.Duration) error {
	task, exists := tasklist.tasks[id]
	if !exists {
		return errors.New("任务不存在")
	}
	task.stateLock.Lock()
	defer task.stateLock.Unlock()
	if task.queueing { //如果已经在排队
		return errors.New("任务已启动")
	}
	Daemon.AddTask(task, duration)
	task.queueing = true
	return nil
}

//将一个任务停止执行
//
//返回任务是否存在和错误信息
func (tasklist *taskList) Stop(id string) error {
	task, exists := tasklist.tasks[id]
	if !exists {
		return errors.New("任务不存在")
	}
	task.stateLock.Lock()
	defer task.stateLock.Unlock()
	if !task.queueing { //如果都没在排队
		return nil //直接返回成功
	}
	if task.IsRunning() { //如果在运行
		return task.Stop() //那就停止
	} else if task.queueing { //如果在排队
		Daemon.CancelTask(id) //那就取消
	}
	task.queueing = false
	return nil
}
