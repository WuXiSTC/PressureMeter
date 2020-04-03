package TaskList

import (
	"errors"
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"sync"
)

type taskList struct {
	tasks map[string]*task
}

var TaskList = taskList{make(map[string]*task)}

//插入一个任务
//
//应该先删除ID对应的任务再插入
func (tasklist *taskList) AddTask(t TaskInterface) error {
	tsk := &task{t, new(sync.Mutex), false}
	_, exists := tasklist.tasks[tsk.GetID()]
	if exists {
		return errors.New("任务已存在")
	}
	tasklist.tasks[tsk.GetID()] = tsk
	return nil
}

//按照ID获取任务
//
//返回任务信息获取接口和是否存在
func (tasklist *taskList) GetInfo(id string) TaskInfo {
	return tasklist.tasks[id]
}

type TaskState int

const (
	NOTEXISTS TaskState = -1
	QUEUEING  TaskState = 0
	RUNNING   TaskState = 1
	STOPPED   TaskState = 2
)

func (tasklist *taskList) GetState(id string) TaskState {
	task, exists := tasklist.tasks[id]
	if !exists {
		return NOTEXISTS
	}
	if task.queueing && !task.IsRunning() {
		return QUEUEING
	}
	if task.queueing && task.IsRunning() {
		return RUNNING
	}
	return STOPPED
}

//按照ID删除任务
//
//返回任务是否存在和错误信息
func (tasklist *taskList) DelTask(id string) (exists bool, err error) {
	tsk, exists := tasklist.tasks[id]
	if exists {
		err = tsk.Delete()
		delete(tasklist.tasks, id)
	}
	return
}

func (tasklist *taskList) Exists(id string) bool {
	_, exists := tasklist.tasks[id]
	return exists
}

//停止所有任务
func (tasklist *taskList) StopAll() error {
	Daemon.Stop()
	for _, tsk := range tasklist.tasks {
		if err := tsk.Stop(); err != nil {
			return err
		}
	}
	util.Log("All tasks stopped")
	return nil
}

//删除所有任务
func (tasklist *taskList) DelAll() error {
	for _, tsk := range tasklist.tasks {
		if err := tsk.Delete(); err != nil {
			return err
		}
	}
	util.Log("All tasks deleted")
	tasklist.tasks = nil
	return nil
}
