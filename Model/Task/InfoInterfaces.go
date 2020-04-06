package Task

import (
	"gitee.com/WuXiSTC/PressureMeter/util"
)

//获取设置文件路径，没有就先创建
func (tsk *task) GetConfigFilePath() string {
	if err := util.MakeFile(tsk.configFilePath); err != nil {
		return ""
	}
	return tsk.configFilePath
}

//获取结果文件路径，没有就先创建
func (tsk *task) GetResultFilePath() string {
	if err := util.MakeFile(tsk.resultFilePath); err != nil {
		return ""
	}
	return tsk.resultFilePath
}

//获取日志文件路径，没有就先创建
func (tsk *task) GetLogFilePath() string {
	if err := util.MakeFile(tsk.logFilePath); err != nil {
		return ""
	}
	return tsk.logFilePath
}
