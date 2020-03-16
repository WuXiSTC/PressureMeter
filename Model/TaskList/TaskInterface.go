package TaskList

import "PressureMeter/Model/Daemon"

const ( //Task的三种状态
	STATE_STOPPED  = iota //停止
	STATE_QUEUEING        //在队列中
	STATE_RUNNING         //正在运行
)

type TaskInfo interface {
	GetConfigFilePath() string
	GetResultFilePath() string
	GetLogFilePath() string
	GetStateCode() int
}

type TaskInterface interface {
	Daemon.TaskInterface
	TaskInfo
	Delete() error
	GetState() int
	SetState(int)
}
