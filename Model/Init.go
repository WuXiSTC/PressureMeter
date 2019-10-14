package Model

import (
	"../util"
	"./Daemon"
	"./Task"
	tasklist "./TaskList"
)

type Config struct {
	DaemonConf Daemon.Config
	TaskConf   Task.Config
}

var Conf = Config{Daemon.Config{TaskQSize: 1000, TaskAccN: 4},
	Task.Config{JmxDir: "Data/jmx", JtlDir: "Data/jtl", LogDir: "Data/log"}}

var TaskList = tasklist.TaskList

//Model层组件初始化
func Init(configPath string) {
	util.GetConf(configPath, Conf)
	Daemon.Init(Conf.DaemonConf)
	Task.Init(Conf.TaskConf)
}
