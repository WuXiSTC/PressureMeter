package Task

import (
	"bufio"
	"context"
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
	if tsk.IsRunning() { //如果已经在运行
		return nil //就退出
	}
	tsk.shutdownPort = shutdownPort                          //获取关机端口
	tsk.command = conf.getStartCommand(tsk.id, shutdownPort) //获取运行指令
	ctx, cancel := context.WithCancel(context.Background())  //新建运行标记
	tsk.ctx = ctx

	if stdout, err := tsk.command.StdoutPipe(); err == nil { //获取输出
		buf := bufio.NewReader(stdout)
		go func() { //命令行输出线程
			for {
				line, _, _ := buf.ReadLine()
				util.Log(fmt.Sprintf("Jmeter(shutdownPort %d)-->", shutdownPort) + string(line))
				select {
				case <-ctx.Done():
					return
				default:
					continue
				}
			}
		}()
	} else {
		util.LogE(err)
	}

	if err := tsk.command.Start(); err != nil { //启动
		cancel()
		tsk.command = nil
		return err //失败即退出
	}
	util.Log(fmt.Sprintf("Task %s started in PID %d", tsk.id, tsk.command.Process.Pid))

	errChan := make(chan error, 1)
	go func() { //等待报错线程
		errChan <- tsk.command.Wait()
		cancel()
	}()
	go func() { //定时退出线程
		select {
		case <-time.After(duration):
		case err := <-errChan:
			util.LogE(err)
		}
		for { //如果退出消息发送太早进程会接收不到，加个for循环确保100%退出
			util.LogE(tsk.Stop())
			select {
			case <-ctx.Done():
				return
			default:
				continue
			}
		}
	}()
	return nil
}

//用于Daemon的接口，等待任务完成，完成后清理资源
func (tsk *task) Wait() {
	<-tsk.ctx.Done()
}

//用于Daemon的接口，停止任务运行
//
//先停止并删除进程再释放文件
func (tsk *task) Stop() error {
	tsk.stateLock.Lock()
	defer tsk.stateLock.Unlock()
	if !tsk.IsRunning() { //如果已经停止就直接成功
		return nil
	}
	if tsk.command.Process != nil {
		stopCmd := conf.getStopCommand(tsk.shutdownPort)
		if err := stopCmd.Start(); err == nil {
			return stopCmd.Wait()
		} else {
			util.LogE(err)
		}
		if err := tsk.command.Process.Signal(syscall.SIGKILL); err != nil { //停止就是向进程发送kill命令
			return err
		}
	}
	return nil
}
