package Model

import (
	"../util"
	"os"
	"path/filepath"
)

type Config struct {
	jmxDir string
	jtlDir string
}

var Conf = Config{"Data/jmx", "Data/jtl"}

func (conf *Config) jmxPath(id string) string {
	util.LogE(os.MkdirAll(Conf.jmxDir, os.ModePerm)) //没有目录先建目录
	return filepath.Join(conf.jmxDir, id) + ".jmx"   //文件名是任务的id
}

func (conf *Config) jtlPath(id string) string {
	util.LogE(os.MkdirAll(Conf.jtlDir, os.ModePerm)) //没有目录先建目录
	return filepath.Join(conf.jtlDir, id) + ".jtl"   //文件名是任务的id
}
