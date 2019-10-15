package TaskList

import (
	"../Daemon"
	"errors"
)

//将一个任务加进任务队列
func (tasklist *taskList) Start(id string) error {
	task, exists := tasklist.tasks[id]
	if !exists {
		return errors.New("任务不存在")
	}
	state := (*task).GetState()
	if state == STATE_QUEUEING || state == STATE_RUNNING {
		return nil
	}
	(*task).SetState(STATE_QUEUEING)
	Daemon.AddTask(*task)
	return nil
}

//将一个任务停止执行
func (tasklist *taskList) Stop(id string) error {
	task, exists := tasklist.tasks[id]
	if !exists {
		return errors.New("任务不存在")
	}
	switch (*task).GetState() {
	case STATE_RUNNING: //如果在运行
		return (*task).Stop() //那就停止
	case STATE_QUEUEING: //如果在队列中
		Daemon.CancelTask(id) //那就取消
		(*task).SetState(STATE_STOPPED)
		return nil
	}
	return nil
}
