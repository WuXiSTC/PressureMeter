package Model

import (
	"../util"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

type task struct {
	id             string
	configFilePath string
	resultFilePath string
}

//新建Task
func Task(id string, configFile multipart.File) (*task, error) {
	err := os.MkdirAll(Conf.jmxPath, os.ModePerm) //没有目录先建目录
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(Conf.jtlPath, os.ModePerm) //没有目录先建目录
	if err != nil {
		return nil, err
	}

	configFilePath := filepath.Join(Conf.jmxPath, id) + ".jmx" //文件名是任务的id
	resultFilePath := filepath.Join(Conf.jtlPath, id) + ".jtl" //文件名是任务的id

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

	jtl, err := os.OpenFile(resultFilePath, os.O_WRONLY|os.O_CREATE, os.ModePerm) //打开文件流
	if err != nil {
		return nil, err
	}
	defer func() {
		util.LogE(jtl.Close())
	}()

	return &task{id, configFilePath, resultFilePath}, nil
}

func (tsk *task) Delete() error {
	_ = os.Remove(tsk.configFilePath) //删除之前的配置文件
	_ = os.Remove(tsk.resultFilePath) //删除之前的结果文件防止发生追加
	return nil
}
