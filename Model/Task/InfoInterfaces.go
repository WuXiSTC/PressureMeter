package Task

import (
	"PressureMeter/util"
	"os"
)

//获取设置文件路径，没有就先创建
func (tsk *task) GetConfigFilePath() string {
	if err := util.MakeFile(*tsk.configFilePath); err != nil {
		return ""
	}
	return *tsk.configFilePath
}

//获取结果文件路径，没有就先创建
func (tsk *task) GetResultFilePath() string {
	if err := util.MakeFile(*tsk.resultFilePath); err != nil {
		return ""
	}
	return *tsk.resultFilePath
}

//获取日志文件路径，没有就先创建
func (tsk *task) GetLogFilePath() string {
	if tsk.logfile != nil {
		f, err := os.OpenFile(*tsk.logFilePath, os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			util.LogE(err)
			return ""
		}
		tsk.command.Stdout = f
		util.LogE(tsk.logfile.Close()) //强制缓冲区写入文件
		tsk.logfile = f                //并且给新的输出流
	} else if err := util.MakeFile(*tsk.resultFilePath); err != nil {
		return ""
	}
	return *tsk.logFilePath
}

func (tsk *task) GetStateCode() int {
	return *tsk.state
}
