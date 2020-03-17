package Model

import (
	"errors"
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	task "gitee.com/WuXiSTC/PressureMeter/Model/Task"
	tasklist "gitee.com/WuXiSTC/PressureMeter/Model/TaskList"
	"mime/multipart"
)

var TaskList = tasklist.TaskList

var StateList = map[int]string{
	tasklist.STATE_STOPPED:  "Stopped",
	tasklist.STATE_QUEUEING: "Queueing",
	tasklist.STATE_RUNNING:  "Running"}

type Config struct {
	DaemonConfig Daemon.Config
	TaskConfig   task.Config
}

func DefaultConfig() Config {
	return Config{
		DaemonConfig: Daemon.DefaultConfig(),
		TaskConfig:   task.DefaultConfig(),
	}
}

//Model层组件初始化
func Init(c Config) {
	Daemon.Init(c.DaemonConfig)
	task.Init(c.TaskConfig)
}

//新建并添加一个新Task
func AddNewTask(taskId string, jmx multipart.File) error {
	if TaskList.Exists(taskId) {
		return errors.New("任务已存在")
	}
	tsk, err := task.Constructor(taskId, jmx)
	if err != nil {
		return err
	}
	return TaskList.AddTask(tsk)
}
