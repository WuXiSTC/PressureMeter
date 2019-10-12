package Model

import (
	"../util"
	"fmt"
)

type taskList struct {
	tasks map[string]*task
}

var TaskList = taskList{make(map[string]*task)}

//插入一个任务
//
//先删除ID对应的任务再插入
func (tasklist *taskList) AddTask(tsk *task) error {
	tasklist.tasks[tsk.id] = tsk
	return nil
}

//按照ID获取任务
func (tasklist *taskList) GetTask(id string) (*task, bool) {
	tsk, exists := tasklist.tasks[id]
	return tsk, exists
}

//按照ID删除任务
//
//如果任务存在就调用任务的自删除然后从列表中删除任务并返回错误信息，不存在就直接返回无错误
func (tasklist *taskList) DelTask(id string) error {
	tsk, exists := tasklist.tasks[id]
	if exists {
		err := tsk.Delete()
		if err != nil {
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

func (tasklist *taskList) StopAll() {
	for id, tsk := range tasklist.tasks {
		util.LogE(tsk.Stop())
		util.Log(fmt.Sprintf("task %s stopped", id))
	}
}

func (tasklist *taskList) DelAll() {
	for id, tsk := range tasklist.tasks {
		util.LogE(tsk.Delete())
		util.Log(fmt.Sprintf("task %s deleted", id))
	}
	tasklist.tasks = nil
}
