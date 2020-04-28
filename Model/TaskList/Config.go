package TaskList

import (
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	"time"
)

type TaskStateList struct {
	TaskStates map[string]Task `json:"TaskStates"`
	AllTasks   []string        `json:"AllTasks"`
}

type Task struct {
	ID        string        `json:"ID"`
	Duration  time.Duration `json:"Duration"`
	StateCode TaskState     `json:"StateCode"`
	State     string        `json:"State"`
	StartTime int64         `json:"StartTime"`
}

var stateDescriptions = map[TaskState]string{
	NOTEXISTS: "NOTEXISTS",
	QUEUEING:  "QUEUEING",
	RUNNING:   "RUNNING",
	STOPPED:   "STOPPED"}

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
var taskStateList = TaskStateList{
	TaskStates: map[string]Task{},
	AllTasks:   []string{},
}

func Init() {
	Daemon.UpdateStateCallback = func(runnings []Daemon.TaskInterface, startTimes []time.Time, queuings []Daemon.TaskInterface) {
		taskStateList.TaskStates = map[string]Task{}
		for i, t := range runnings {
			if t != nil {
				taskStateList.TaskStates[t.GetID()] = Task{
					ID:        t.GetID(),
					Duration:  t.(TaskInterface).GetDuration(),
					StateCode: RUNNING,
					State:     stateDescriptions[RUNNING],
					StartTime: startTimes[i].UnixNano(),
				}
			}
		}
		for i, t := range queuings {
			taskStateList.TaskStates[t.GetID()] = Task{
				ID:        t.GetID(),
				Duration:  t.(TaskInterface).GetDuration(),
				StateCode: QUEUEING,
				State:     stateDescriptions[QUEUEING],
				StartTime: startTimes[i].UnixNano(),
			}
		}
		UpdateStateCallback(taskStateList)
	}
}
