package TaskList

import (
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	"sync"
)

type TaskInfo interface {
	GetConfigFilePath() string
	GetResultFilePath() string
	GetLogFilePath() string
	IsRunning() bool
}

type TaskInterface interface {
	Daemon.TaskInterface
	TaskInfo
	Delete() error
}

type task struct {
	TaskInterface
	stateLock *sync.Mutex
	queueing  bool
}
