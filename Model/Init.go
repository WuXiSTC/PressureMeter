package Model

import (
	"../util"
	"./Daemon"
	task "./Task"
	tasklist "./TaskList"
	"errors"
	"mime/multipart"
)

type Config struct {
	DaemonConf Daemon.Config `yaml:",inline"`
	TaskConf   task.Config   `yaml:",inline"`
}

var Conf = Config{Daemon.Config{TaskQSize: 1000, TaskAccN: 4},
	task.Config{JmxDir: "Data/jmx", JtlDir: "Data/jtl", LogDir: "Data/log"}}

var TaskList = tasklist.TaskList

var Task = task.Constructor

//Model层组件初始化
func Init(configPath string) {
	util.GetConf(configPath, &Conf)
	Daemon.Init(Conf.DaemonConf)
	task.Init(Conf.TaskConf)
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
