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

func (tsk *task) Stop() error {
	if err := tsk.command.Process.Kill(); err != nil {
		return err
	}
	util.LogE(tsk.logfile.Close())
	tsk.logfile = nil
	return nil
}
