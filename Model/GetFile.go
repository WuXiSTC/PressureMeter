package Model

import (
	"../util"
	"os"
)

func (tsk *task) GetConfigFileStream() *os.File {
	out, err := os.OpenFile(tsk.configFilePath, os.O_RDONLY, os.ModePerm) //打开文件流
	util.LogE(err)
	return out
}

func (tsk *task) GetConfigFilePath() string {
	f, err := os.OpenFile(tsk.configFilePath, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		util.LogE(err)
		return ""
	}
	defer func() { util.LogE(f.Close()) }()
	return tsk.configFilePath
}

func (tsk *task) GetResultFilePath() string {
	f, err := os.OpenFile(tsk.resultFilePath, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		util.LogE(err)
		return ""
	}
	defer func() { util.LogE(f.Close()) }()
	return tsk.resultFilePath
}

func (tsk *task) GetLogFilePath() string {
	f, err := os.OpenFile(tsk.logFilePath, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		util.LogE(err)
		return ""
	}
	defer func() { util.LogE(f.Close()) }()
	return tsk.logFilePath
}
