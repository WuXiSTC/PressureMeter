package TaskList

import (
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	"time"
)

type TaskStateList struct {
	RunningTasks  []RunningTask  `json:"RunningTasks"`
	QueueingTasks []QueueingTask `json:"QueueingTasks"`
	AllTasks      []string       `json:"AllTasks"`
}

type RunningTask struct {
	QueueingTask
	StartTime int64 `json:"StartTime"`
}

type QueueingTask struct {
	ID       string        `json:"ID"`
	Duration time.Duration `json:"Duration"`
}

//当有某个任务的状态发生变化时调用此函数
var UpdateStateCallback func(TaskStateList)

//记录系统中所有任务的状态
var taskStateList TaskStateList

func Init() {
	Daemon.UpdateStateCallback = func(runnings []Daemon.TaskInterface, startTimes []time.Time, queuings []Daemon.TaskInterface) {
		taskStateList.RunningTasks = make([]RunningTask, len(runnings))
		for i, t := range runnings {
			if t != nil {
				taskStateList.RunningTasks[i] = RunningTask{
					QueueingTask: QueueingTask{
						ID:       t.GetID(),
						Duration: t.(TaskInterface).GetDuration(),
					},
					StartTime: startTimes[i].UnixNano(),
				}
			}
		}
		taskStateList.QueueingTasks = make([]QueueingTask, len(queuings))
		for i, t := range queuings {
			taskStateList.QueueingTasks[i] = QueueingTask{
				ID:       t.GetID(),
				Duration: t.(TaskInterface).GetDuration(),
			}
		}
		UpdateStateCallback(taskStateList)
	}
}
