package Task

import (
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"syscall"
	"time"
)

func (tsk *task) GetID() string {
	return tsk.id
}

//用于Daemon的接口，开始任务执行
func (tsk *task) Start(shutdownPort uint16, duration time.Duration) error {
	tsk.stateLock.Lock()
	defer tsk.stateLock.Unlock()
	if tsk.running { //如果已经在运行
		return nil //就退出
	}
	tsk.shutdownPort = shutdownPort                          //获取关机端口
	tsk.command = conf.getStartCommand(tsk.id, shutdownPort) //获取运行指令
	if err := tsk.command.Start(); err != nil {              //启动
		return err //失败即退出
	}
	util.Log(fmt.Sprintf("Task %s started in PID %d", tsk.id, tsk.command.Process.Pid))
	tsk.running = true

	go func() {
		errChan := make(chan error, 1)
		go func() {
			errChan <- tsk.command.Wait()
		}()
		select {
		case <-time.After(duration):
		case err := <-errChan:
			util.LogE(err)
		}
		util.LogE(tsk.Stop())
	}()
	return nil
}

//用于Daemon的接口，等待任务完成，完成后清理资源
func (tsk *task) Wait() error {
	util.LogE(tsk.command.Wait())
	tsk.stateLock.Lock()
	defer tsk.stateLock.Unlock()
	tsk.command = nil //进程完成后清除进程
	tsk.running = false
	return nil
}

//用于Daemon的接口，停止任务运行
//
//先停止并删除进程再释放文件
func (tsk *task) Stop() error {
	tsk.stateLock.Lock()
	defer tsk.stateLock.Unlock()
	if !tsk.running { //如果已经停止就直接成功
		return nil
	}
	if tsk.command.Process != nil {
		if err := conf.getStopCommand(tsk.shutdownPort).Start(); err == nil {
			return nil
		} else {
			util.LogE(err)
		}
		if err := tsk.command.Process.Signal(syscall.SIGKILL); err != nil { //停止就是向进程发送kill命令
			return err
		}
	} else {
		tsk.running = false
	}
	return nil
}
