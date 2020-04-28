package Model

import (
	"errors"
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	task "gitee.com/WuXiSTC/PressureMeter/Model/Task"
	tasklist "gitee.com/WuXiSTC/PressureMeter/Model/TaskList"
	"mime/multipart"
)

var TaskList = tasklist.TaskList
var Tasks tasklist.TaskStateList

type Config struct {
	DaemonConfig Daemon.Config `yaml:"DaemonConfig" usage:"Configuration of PressureMater Daemon."`
	TaskConfig   task.Config   `yaml:"TaskConfig" usage:"Configuration of PressureMater Task."`
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
	task.Init(c.TaskConfig, c.DaemonConfig)
	tasklist.UpdateStateCallback = func(list tasklist.TaskStateList) {
		Tasks = list
	}
	tasklist.Init()
}

//新建并添加一个新Task
func AddNewTask(taskId string, jmx multipart.File) error {
	if TaskList.GetState(taskId) != tasklist.NOTEXISTS {
		return errors.New("任务已存在")
	}
	tsk, err := task.New(taskId, jmx)
	if err != nil {
		return err
	}
	return TaskList.AddTask(tsk)
}
