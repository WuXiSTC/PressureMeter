package Task

import (
	"context"
	"errors"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"sync"
	"time"
)

type task struct {
	id             string //任务的ID
	configFilePath string //设置文件的路径
	resultFilePath string //结果文件的路径
	logFilePath    string //日志文件的路径
	stateLock      *sync.RWMutex
	ctx            context.Context
	duration       time.Duration
}

//新建Task，相当于Task的构造函数
//
//输入值是任务ID和配置文件流，返回新建的任务指针和错误信息
func New(id string, configFile multipart.File) (*task, error) {
	configFilePath := conf.jmxPath(id) //文件名是任务的id
	resultFilePath := conf.jtlPath(id) //文件名是任务的id
	logFilePath := conf.logPath(id)    //文件名是任务的id

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

	if err = util.EmptyFile(resultFilePath); err != nil {
		return nil, err //创建结果文件
	}
	if err = util.EmptyFile(logFilePath); err != nil {
		return nil, err //创建日志文件
	}

	ctx, cancel := context.WithCancel(context.Background())
	tsk := &task{id, configFilePath, resultFilePath, logFilePath,
		new(sync.RWMutex), ctx, 0}
	cancel()
	return tsk, nil
}

//删除任务，顺带删除任务相关文件
//
//先停止任务，然后删除任务相关文件
func (tsk *task) Delete() error {
	if tsk.isRunning() {
		return errors.New("任务未停止，无法删除")
	}
	if err := util.DeleteFile(tsk.configFilePath); err != nil { //删除之前的配置文件
		return err
	}
	if err := util.DeleteFile(tsk.resultFilePath); err != nil { //删除之前的结果文件防止发生追加
		return err
	}
	if err := util.DeleteFile(tsk.logFilePath); err != nil { //删除之前的日志文件防止发生追加
		return err
	}
	return nil
}

func (tsk *task) SetDuration(d time.Duration) {
	tsk.duration = d
}

func (tsk task) GetDuration() time.Duration {
	return tsk.duration
}
