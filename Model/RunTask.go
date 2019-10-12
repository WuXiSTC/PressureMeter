package Model

import (
	"../util"
	"os"
)

func (tsk *task) Start() error {
	f, err := os.OpenFile(tsk.logFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	tsk.logfile = f
	tsk.command.Stdout = f
	if err := tsk.command.Start(); err != nil {
		util.LogE(f.Close())
		return err
	}
	return nil
}

//停止任务运行
//
//先停止并删除进程再释放文件
func (tsk *task) Stop() error {
	if tsk.command.Process != nil {
		if err := tsk.command.Process.Kill(); err != nil {
			return err
		}
		tsk.command.Process = nil
	}
	util.LogE(tsk.logfile.Close())
	tsk.logfile = nil
	util.Log("task " + tsk.id + " stopped")
	return nil
}
