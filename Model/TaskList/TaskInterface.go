package TaskList

import (
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	"time"
)

type TaskInfo interface {
	GetConfigFilePath() string
	GetResultFilePath() string
	GetLogFilePath() string
}

type TaskInterface interface {
	Daemon.TaskInterface
	TaskInfo
	Delete() error
	SetDuration(d time.Duration)
	GetDuration() time.Duration
}
