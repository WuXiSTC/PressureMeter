package TaskList

import (
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
)

//将一个任务加进任务队列
//
//不会返回错误，返回的bool表示任务是否存在
func (tasklist *taskList) Start(id string) bool {
	task, exists := tasklist.tasks[id]
	if !exists {
		return exists
	}
	state := (*task).GetState()
	if state == STATE_QUEUEING || state == STATE_RUNNING {
		return exists
	}
	(*task).SetState(STATE_QUEUEING)
	Daemon.AddTask(*task)
	return exists
}

//将一个任务停止执行
//
//返回任务是否存在和错误信息
func (tasklist *taskList) Stop(id string) (bool, error) {
	task, exists := tasklist.tasks[id]
	if !exists {
		return exists, nil
	}
	switch (*task).GetState() {
	case STATE_RUNNING: //如果在运行
		return exists, (*task).Stop() //那就停止
	case STATE_QUEUEING: //如果在队列中
		Daemon.CancelTask(id) //那就取消
		(*task).SetState(STATE_STOPPED)
		return exists, nil
	}
	return exists, nil
}
