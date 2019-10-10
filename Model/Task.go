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
	err := os.Mkdir(Conf.jmxPath, os.ModePerm) //没有目录先建目录
	if err == nil {
		util.Log("jmx dir " + Conf.jmxPath + " created")
	}

	pathPrefix := filepath.Join(Conf.jmxPath, id) //文件名是任务的id
	configFilePath := pathPrefix + ".jmx"
	resultFilePath := pathPrefix + ".jtl"

	out, err := os.OpenFile(configFilePath, os.O_WRONLY|os.O_CREATE, os.ModePerm) //打开文件流
	if err != nil {
		return nil, err
	}
	defer func() {
		util.LogE(out.Close())
	}()

	n, err := io.Copy(out, configFile) //写入配置文件
	if err != nil {
		return nil, err
	}
	util.Log(strconv.FormatInt(n, 10) + " byte jmx received, saved to " + configFilePath)

	_ = os.Remove(resultFilePath) //删除之前的结果文件防止发生追加

	return &task{id, configFilePath, resultFilePath}, nil
}

func (task *task) Delete() error {
	_ = os.Remove(task.configFilePath) //删除之前的配置文件
	_ = os.Remove(task.resultFilePath) //删除之前的结果文件防止发生追加
	return nil
}
