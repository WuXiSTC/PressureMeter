package TaskList

import (
	"errors"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"sync"
)

var tasklist map[string]TaskInterface
var tasklistMu = new(sync.RWMutex)

//插入一个任务
//
//应该先删除ID对应的任务再插入
func AddTask(tsk TaskInterface) error {
	tasklistMu.Lock()
	defer tasklistMu.Unlock()
	_, exists := tasklist[tsk.GetID()]
	if exists {
		return errors.New("任务已存在")
	}
	tasklist[tsk.GetID()] = tsk
	return nil
}

//按照ID获取任务
//
//返回任务信息获取接口和是否存在
func GetInfo(id string) TaskInfo {
	tasklistMu.RLock()
	defer tasklistMu.RUnlock()
	if info, exists := tasklist[id]; exists {
		return info
	}
	return nil
}

//按照ID删除任务
//
//返回任务是否存在和错误信息
func DelTask(id string) (err error) {
	tasklistMu.Lock()
	defer tasklistMu.Unlock()
	switch getState(id) {
	case STOPPED: //存在且已停止，可删
		if tsk, exists := tasklist[id]; exists {
			if err = tsk.Delete(); err == nil {
				delete(tasklist, id)
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
func DelAll() error {
	tasklistMu.Lock()
	defer tasklistMu.Unlock()
	StopAll()
	for id := range tasklist {
		if err := DelTask(id); err != nil {
			return err
		}
	}
	util.Log("All tasks deleted")
	tasklist = nil
	return nil
}
