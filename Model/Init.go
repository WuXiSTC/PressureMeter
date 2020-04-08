package Model

import (
	"errors"
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	task "gitee.com/WuXiSTC/PressureMeter/Model/Task"
	tasklist "gitee.com/WuXiSTC/PressureMeter/Model/TaskList"
	"mime/multipart"
)

var TaskList = tasklist.TaskList

type Config struct {
	DaemonConfig Daemon.Config `yaml:"DaemonConfig" usage:"Configuration of PressureMater Daemon."`
	TaskConfig   task.Config   `yaml:"TaskConfig" usage:"Configuration of PressureMater Task."`

	//当有某个任务的状态发生变化时此函数将被调用
	UpdateStateCallback func(list tasklist.TaskStateList) `yaml:"-"`
}

func DefaultConfig() Config {
	return Config{
		DaemonConfig:        Daemon.DefaultConfig(),
		TaskConfig:          task.DefaultConfig(),
		UpdateStateCallback: func(tasklist.TaskStateList) {},
	}
}

//Model层组件初始化
func Init(c Config) {
	Daemon.Init(c.DaemonConfig)
	task.Init(c.TaskConfig, c.DaemonConfig)
	tasklist.UpdateStateCallback = c.UpdateStateCallback
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
