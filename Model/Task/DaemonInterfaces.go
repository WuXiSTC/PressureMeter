package Task

import (
	"bufio"
	"context"
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"regexp"
	"time"
)

//此任务是否在运行，非线程安全，需要额外保护
func (tsk *task) isRunning() bool {
	select {
	case <-tsk.ctx.Done(): //如果已经结束
		return false //就是不在运行
	default: //否则
		return true //就是在运行
	}
}

//用于Daemon的接口，获取任务ID
func (tsk *task) GetID() string {
	return tsk.id
}

var rex, _ = regexp.Compile("\\s")

//用于Daemon的接口，开始任务执行
func (tsk *task) Start(i uint16) {
	tsk.stateLock.Lock()
	defer tsk.stateLock.Unlock()
	if tsk.isRunning() { //如果已经在运行
		return //就退出
	}
	command := conf.startCommand(tsk.id, i)                 //获取运行指令
	ctx, cancel := context.WithCancel(context.Background()) //新建运行标记
	tsk.ctx = ctx

	if stdout, err := command.StdoutPipe(); err == nil { //获取输出
		buf := bufio.NewReader(stdout)
		go func() { //命令行输出线程
			for {
				line, _, _ := buf.ReadLine()
				s := string(line)
				if len(rex.ReplaceAllString(s, "")) > 0 {
					util.Log(fmt.Sprintf("Jmeter log(thread %d)-->", i) + s)
				}
				if !tsk.isRunning() { //如果不在运行
					return //就退出
				}
				time.Sleep(5e8)
			}
		}()
	} else {
		util.LogE(err)
	}

	if err := command.Start(); err != nil { //启动
		cancel()
		command = nil
		return //失败即退出
	}
	util.Log(fmt.Sprintf("Task %s started in PID %d", tsk.id, command.Process.Pid))

	errChan := make(chan error, 1)
	go func() { //等待报错线程
		errChan <- command.Wait()
		cancel()
	}()
	go func(duration time.Duration) {
		select {
		case <-time.After(duration): //定时退出
			tsk.Stop(i)
		case err := <-errChan: //自然退出
			util.LogE(err)
		}
	}(tsk.duration)
}

//用于Daemon的接口，等待任务完成，完成后清理资源
func (tsk *task) Wait() {
	tsk.stateLock.RLock()
	ctx := tsk.ctx
	tsk.stateLock.RUnlock()
	<-ctx.Done()
}

//用于Daemon的接口，停止任务运行，持续发送退出消息直到退出
func (tsk *task) Stop(i uint16) {
	tsk.stateLock.Lock()
	defer tsk.stateLock.Unlock()
	for tsk.isRunning() { //如果退出消息发送太早进程会接收不到，加个for循环确保100%退出
		stopCmd := conf.stopCommand(i)
		util.LogE(stopCmd.Start())
		util.LogE(stopCmd.Wait())
		time.Sleep(1e8)
	}
}
