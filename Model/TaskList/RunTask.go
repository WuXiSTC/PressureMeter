package TaskList

import (
	"errors"
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"sync"
	"time"
)

var daemonMu = new(sync.RWMutex)

//将一个任务加进任务队列
//
//不会返回错误，返回任务是否存在
func Start(id string, duration time.Duration) error {
	tasklistMu.RLock()
	defer tasklistMu.RUnlock()
	daemonMu.Lock()
	defer daemonMu.Unlock()
	switch getState(id) {
	case NOTEXISTS:
		return errors.New("not exists")
	case STOPPED:
		task := tasklist[id]
		task.SetDuration(duration)
		Daemon.Queue(task)
	default:
		return errors.New("already started")
	}
	return nil
}

//将一个任务停止执行
//
//返回任务是否存在
func Stop(id string) error {
	tasklistMu.RLock()
	defer tasklistMu.RUnlock()
	daemonMu.Lock()
	defer daemonMu.Unlock()
	switch getState(id) {
	case NOTEXISTS:
		return errors.New("not exists")
	case STOPPED:
		return errors.New("already stopped")
	default:
		Daemon.Cancel(id)
		return nil
	}
}

type TaskState int

const (
	NOTEXISTS TaskState = -1
	QUEUEING  TaskState = 0
	RUNNING   TaskState = 1
	STOPPED   TaskState = 2
)

func GetState(id string) TaskState {
	tasklistMu.RLock()
	defer tasklistMu.RUnlock()
	daemonMu.RLock()
	defer daemonMu.RUnlock()
	return getState(id)
}

func getState(id string) TaskState {
	switch Daemon.GetState(id) {
	case Daemon.RUNNING:
		return RUNNING
	case Daemon.QUEUEING:
		return QUEUEING
	default:
		if _, exists := tasklist[id]; !exists {
			return NOTEXISTS
		} else {
			return STOPPED
		}
	}
}

//停止所有任务
func StopAll() {
	tasklistMu.RLock()
	defer tasklistMu.RUnlock()
	daemonMu.Lock()
	defer daemonMu.Unlock()
	Daemon.Stop()
	for id := range tasklist {
		Daemon.Cancel(id)
	}
	util.Log("All tasks stopped")
}
