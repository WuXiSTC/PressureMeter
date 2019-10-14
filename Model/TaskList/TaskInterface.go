package TaskList

import "../Daemon"

const ( //Task的三种状态
	STATE_STOPPED  = iota //停止
	STATE_QUEUEING        //在队列中
	STATE_RUNNING         //正在运行
)

type TaskInterface interface {
	Daemon.TaskInterface
	Stop() error
	Delete() error
	GetState() int
	SetState(int)
}
