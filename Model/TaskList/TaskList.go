package TaskList

import (
	"errors"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"sync"
)

var TaskList = taskList{
	list: map[string]TaskInterface{},
	mu:   new(sync.RWMutex),
}

type taskList struct {
	list map[string]TaskInterface
	mu   *sync.RWMutex
}

//插入一个任务
//
//应该先删除ID对应的任务再插入
func (tl *taskList) AddTask(tsk TaskInterface) error {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	_, exists := tl.list[tsk.GetID()]
	if exists {
		return errors.New("任务已存在")
	}
	tl.list[tsk.GetID()] = tsk
	return nil
}

//按照ID获取任务
//
//返回任务信息获取接口和是否存在
func (tl *taskList) GetInfo(id string) TaskInfo {
	tl.mu.RLock()
	defer tl.mu.RUnlock()
	if info, exists := tl.list[id]; exists {
		return info
	}
	return nil
}

//按照ID删除任务
//
//返回任务是否存在和错误信息
func (tl *taskList) DelTask(id string) (err error) {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	switch tl.getState(id) {
	case STOPPED: //存在且已停止，可删
		if tsk, exists := tl.list[id]; exists {
			if err = tsk.Delete(); err == nil {
				delete(tl.list, id)
			}
		}
	case NOTEXISTS: //不存在
		err = errors.New("not exists")
	default:
		err = errors.New("running or queuing")
	}
	return
}

//删除所有任务
func (tl *taskList) DelAll() error {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	tl.StopAll()
	for id := range tl.list {
		if err := tl.DelTask(id); err != nil {
			return err
		}
	}
	util.Log("All tasks deleted")
	tl.list = nil
	return nil
}
