package Task

import (
	"bufio"
	"errors"
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/Model/TaskList"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"os"
	"regexp"
	"strconv"
	"syscall"
)

func (tsk *task) GetID() string {
	return *tsk.id
}

//用于Daemon的接口，开始任务执行
func (tsk *task) Start() error {
	f, err := os.OpenFile(*tsk.logFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	tsk.logfile = f
	tsk.command.Stdout = f
	if err := tsk.command.Start(); err != nil {
		util.LogE(f.Close())
		return err
	}
	util.Log(fmt.Sprintf("Task %d started in PID %d", tsk.id, tsk.command.Process.Pid))
	tsk.SetState(TaskList.STATE_RUNNING)
	return nil
}

//用于Daemon的接口，等待任务完成，完成后清理资源
func (tsk *task) Wait() error {
	util.LogE(tsk.command.Wait())
	util.LogE(tsk.logfile.Close())
	tsk.command.Stdout = nil
	tsk.logfile = nil
	tsk.command = conf.getCommand(*tsk.id) //进程完成后重开进程
	tsk.SetState(TaskList.STATE_STOPPED)
	return nil
}

//用于Daemon的接口，停止任务运行
//
//先停止并删除进程再释放文件
func (tsk *task) Stop() error {
	if tsk.GetState() == TaskList.STATE_STOPPED { //如果已经停止就直接成功
		return nil
	}
	if tsk.command.Process != nil {
		if err := tsk.sendStopMsg(); err == nil {
			return nil
		} else {
			util.LogE(err)
		}
		if err := tsk.command.Process.Signal(syscall.SIGKILL); err != nil { //停止就是向进程发送kill命令
			return err
		}
	} else {
		tsk.SetState(TaskList.STATE_STOPPED)
	}
	return nil
}

//向任务停止监听端口发送停止信号
func (tsk *task) sendStopMsg() error {
	port, err := tsk.getStopPort()
	if err != nil {
		return err
	}
	stopCommand := getStopCommand(port)
	return stopCommand.Start()
}

//获取停止信号的监听端口，以读log的形式搞的，肥肠暴力
func (tsk *task) getStopPort() (int, error) {
	logPath := tsk.GetLogFilePath()
	f, err := os.OpenFile(logPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer func() { util.LogE(f.Close()) }()

	re, _ := regexp.Compile("Waiting for possible Shutdown/StopTestNow/HeapDump/ThreadDump message on port (.*)$")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for i := 0; i < 10 && scanner.Scan(); i++ {
		s := re.FindSubmatch([]byte(scanner.Text()))
		if len(s) >= 2 {
			portStr := string(s[1])
			port, err := strconv.Atoi(portStr)
			if err != nil {
				util.Log(fmt.Sprintf("Fail to detect port: %s", portStr))
				return 0, err
			}
			return port, nil
		}
	}
	return 0, errors.New("找不到端口记录")
}
