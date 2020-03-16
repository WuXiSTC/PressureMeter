package TaskList

import (
	"errors"
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	"gitee.com/WuXiSTC/PressureMeter/util"
)

type taskList struct {
	tasks map[string]*TaskInterface
}

var TaskList = taskList{make(map[string]*TaskInterface)}

//插入一个任务
//
//应该先删除ID对应的任务再插入
func (tasklist *taskList) AddTask(tsk TaskInterface) error {
	_, exists := tasklist.tasks[tsk.GetID()]
	if exists {
		return errors.New("任务已存在")
	}
	tasklist.tasks[tsk.GetID()] = &tsk
	return nil
}

//按照ID获取任务
//
//返回任务信息获取接口和是否存在
func (tasklist *taskList) GetInfo(id string) (*TaskInfo, bool) {
	tsk, exists := tasklist.tasks[id]
	if tsk != nil {
		var tskI TaskInfo = *tsk
		return &tskI, exists
	}
	return nil, exists
}

//按照ID删除任务
//
//返回任务是否存在和错误信息
func (tasklist *taskList) DelTask(id string) (bool, error) {
	tsk, exists := tasklist.tasks[id]
	if exists {
		if err := (*tsk).Delete(); err != nil {
			return exists, err
		}
		delete(tasklist.tasks, id)
	}
	return exists, nil
}

func (tasklist *taskList) Exists(id string) bool {
	_, exists := tasklist.tasks[id]
	return exists
}

//停止所有任务
func (tasklist *taskList) StopAll() error {
	Daemon.Stop()
	for _, tsk := range tasklist.tasks {
		if err := (*tsk).Stop(); err != nil {
			return err
		}
	}
	util.Log("All tasks stopped")
	return nil
}

//删除所有任务
func (tasklist *taskList) DelAll() error {
	for _, tsk := range tasklist.tasks {
		if err := (*tsk).Delete(); err != nil {
			return err
		}
	}
	util.Log("All tasks deleted")
	tasklist.tasks = nil
	return nil
}
