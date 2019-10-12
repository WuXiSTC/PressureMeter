package Model

import (
	"../util"
	"os"
)

func (tsk *task) Start() error {
	f, err := os.OpenFile(tsk.logFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	tsk.logfile = f
	tsk.command.Stdout = f
	if err := tsk.command.Start(); err != nil {
		util.LogE(f.Close())
		return err
	}
	tsk.state = STATE_RUNNING
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
	tsk.command.Stdout = nil
	tsk.logfile = nil
	tsk.state = STATE_STOPPED
	util.Log("task " + tsk.id + " stopped")
	return nil
}
