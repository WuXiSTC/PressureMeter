package Task

import (
	"../../util"
	"os"
)

//获取设置文件路径，没有就先创建
func (tsk *task) GetConfigFilePath() string {
	f, err := os.OpenFile(tsk.configFilePath, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		util.LogE(err)
		return ""
	}
	defer func() { util.LogE(f.Close()) }()
	return tsk.configFilePath
}

//获取结果文件路径，没有就先创建
func (tsk *task) GetResultFilePath() string {
	f, err := os.OpenFile(tsk.resultFilePath, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		util.LogE(err)
		return ""
	}
	defer func() { util.LogE(f.Close()) }()
	return tsk.resultFilePath
}

//获取日志文件路径，没有就先创建
func (tsk *task) GetLogFilePath() string {
	f, err := os.OpenFile(tsk.logFilePath, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		util.LogE(err)
		return ""
	}
	tsk.command.Stdout = f
	util.LogE(tsk.logfile.Close()) //强制缓冲区写入文件
	tsk.logfile = f                //并且给新的输出流
	return tsk.logFilePath
}
