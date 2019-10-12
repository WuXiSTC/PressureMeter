package Model

import (
	"../util"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"strconv"
)

type task struct {
	id             string
	configFilePath string
	resultFilePath string
	logFilePath    string
	command        *exec.Cmd
	logfile        *os.File
}

//创建一个新的空文件，或者清空文件
func createFile(path string) error {
	_ = os.Remove(path)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm) //打开文件流
	if err != nil {
		return err
	}
	util.LogE(f.Close())
	return nil
}

//新建Task
func Task(id string, configFile multipart.File) (*task, error) {
	configFilePath := Conf.jmxPath(id) //文件名是任务的id
	resultFilePath := Conf.jtlPath(id) //文件名是任务的id
	logFilePath := Conf.logPath(id)    //文件名是任务的id

	jmx, err := os.OpenFile(configFilePath, os.O_WRONLY|os.O_CREATE, os.ModePerm) //打开文件流
	if err != nil {
		return nil, err
	}
	defer func() {
		util.LogE(jmx.Close())
	}()

	n, err := io.Copy(jmx, configFile) //写入配置文件
	if err != nil {
		return nil, err
	}
	util.Log(strconv.FormatInt(n, 10) + " byte jmx received, saved to " + configFilePath)

	if err = createFile(resultFilePath); err != nil {
		return nil, err
	}
	if err = createFile(logFilePath); err != nil {
		return nil, err
	}

	return &task{id, configFilePath, resultFilePath, logFilePath,
		Conf.getCommand(id), nil}, nil
}

func (tsk *task) Delete() error {
	_ = os.Remove(tsk.configFilePath) //删除之前的配置文件
	_ = os.Remove(tsk.resultFilePath) //删除之前的结果文件防止发生追加
	_ = os.Remove(tsk.logFilePath)    //删除之前的日志文件防止发生追加
	return nil
}

//TODO:删除任务要先检查停止状态
