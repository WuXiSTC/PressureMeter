package Model

import (
	"../util"
	"./Daemon"
	task "./Task"
	tasklist "./TaskList"
)

type Config struct {
	DaemonConf Daemon.Config
	TaskConf   task.Config
}

var Conf = Config{Daemon.Config{TaskQSize: 1000, TaskAccN: 4},
	task.Config{JmxDir: "Data/jmx", JtlDir: "Data/jtl", LogDir: "Data/log"}}

var TaskList = tasklist.TaskList

var Task = task.Constructor

//Model层组件初始化
func Init(configPath string) {
	util.GetConf(configPath, Conf)
	Daemon.Init(Conf.DaemonConf)
	task.Init(Conf.TaskConf)
}
