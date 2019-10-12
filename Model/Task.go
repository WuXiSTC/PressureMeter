package Model

import (
	"../util"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"strconv"
)

const (
	STATE_STOPPED = iota
	STATE_QUEUEING
	STATE_RUNNING
)

type task struct {
	id             string    //任务的ID
	configFilePath string    //设置文件的路径
	resultFilePath string    //结果文件的路径
	logFilePath    string    //日志文件的路径
	command        *exec.Cmd //要执行的指令
	logfile        *os.File  //日志文件流
	state          int       //运行状态
}

//创建一个新的空文件，或者清空文件
func createFile(path string) error {
	_ = os.Remove(path)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm) //打开文件流
	if err != nil {
		return err
	}
	util.LogE(f.Close())
	return nil
}

//新建Task，相当于Task的构造函数
//
//输入值是任务ID和配置文件流，返回新建的任务指针和错误信息
func Task(id string, configFile multipart.File) (*task, error) {
	configFilePath := Conf.jmxPath(id) //文件名是任务的id
	resultFilePath := Conf.jtlPath(id) //文件名是任务的id
	logFilePath := Conf.logPath(id)    //文件名是任务的id

	jmx, err := os.OpenFile(configFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm) //打开文件流
	if err != nil {
		return nil, err
	}
	defer func() {
		util.LogE(jmx.Close()) //结束时关闭
	}()

	n, err := io.Copy(jmx, configFile) //写入配置文件
	if err != nil {
		return nil, err
	}
	util.Log(strconv.FormatInt(n, 10) + " byte jmx received, saved to " + configFilePath)

	if err = createFile(resultFilePath); err != nil {
		return nil, err //创建结果文件
	}
	if err = createFile(logFilePath); err != nil {
		return nil, err //创建日志文件
	}

	return &task{id, configFilePath, resultFilePath, logFilePath,
		Conf.getCommand(id), nil, STATE_STOPPED}, nil
}

//删除任务，顺带删除任务相关文件
//
//先停止任务，然后删除任务相关文件
func (tsk *task) Delete() error {
	if tsk.state == STATE_RUNNING {
		return errors.New("任务正在运行，无法删除")
	}
	_ = os.Remove(tsk.configFilePath) //删除之前的配置文件
	_ = os.Remove(tsk.resultFilePath) //删除之前的结果文件防止发生追加
	_ = os.Remove(tsk.logFilePath)    //删除之前的日志文件防止发生追加
	return nil
}
