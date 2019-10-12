package Model

import (
	"../util"
	"errors"
)

type taskList struct {
	tasks map[string]*task
}

var TaskList = taskList{make(map[string]*task)}

//插入一个任务
//
//应该先删除ID对应的任务再插入
func (tasklist *taskList) AddTask(tsk *task) error {
	_, exists := tasklist.tasks[tsk.id]
	if exists {
		return errors.New("任务已存在")
	}
	tasklist.tasks[tsk.id] = tsk
	return nil
}

//按照ID获取任务
func (tasklist *taskList) GetTask(id string) (*task, bool) {
	tsk, exists := tasklist.tasks[id]
	return tsk, exists
}

//按照ID删除任务
func (tasklist *taskList) DelTask(id string) error {
	tsk, exists := tasklist.tasks[id]
	if exists {
		if err := tsk.Delete(); err != nil {
			return err
		}
		delete(tasklist.tasks, id)
	}
	return nil
}

func (tasklist *taskList) Exists(id string) bool {
	_, exists := tasklist.tasks[id]
	return exists
}

//停止所有任务
func (tasklist *taskList) StopAll() error {
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

//TODO:队列式执行任务
