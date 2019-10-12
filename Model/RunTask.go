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

	go func() { //新开一个线程等待完成，并在线程完成后清理资源
		util.LogE(tsk.command.Wait())
		util.LogE(tsk.logfile.Close())
		tsk.command.Stdout = nil
		tsk.logfile = nil
		tsk.command = Conf.getCommand(tsk.id) //进程完成后重开进程
		tsk.state = STATE_STOPPED
		util.Log("task " + tsk.id + " stopped")
	}()
	return nil
}

//停止任务运行
//
//先停止并删除进程再释放文件
func (tsk *task) Stop() error {
	if tsk.state == STATE_STOPPED { //如果已经停止就直接成功
		return nil
	}
	if tsk.command.Process != nil {
		if err := tsk.command.Process.Kill(); err != nil { //停止就是向进程发送kill命令
			return err
		}
	}
	return nil
}
