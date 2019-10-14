package Task

import (
	"../../util"
	"os"
	"os/exec"
	"path/filepath"
)

type Config struct {
	JmxDir string `yaml:"JmxDir"` //存放jmx文件的目录位置
	JtlDir string `yaml:"JtlDir"` //存放jtl结果文件的目录位置
	LogDir string `yaml:"logDir"` //存放日志文件的目录位置
}

var Conf = Config{"Data/jmx", "Data/jtl", "Data/log"}

func (conf *Config) jmxPath(id string) string {
	util.LogE(os.MkdirAll(Conf.JmxDir, os.ModePerm)) //没有目录先建目录
	return filepath.Join(conf.JmxDir, id) + ".jmx"   //文件名是任务的id
}

func (conf *Config) jtlPath(id string) string {
	util.LogE(os.MkdirAll(Conf.JtlDir, os.ModePerm)) //没有目录先建目录
	return filepath.Join(conf.JtlDir, id) + ".jtl"   //文件名是任务的id
}

func (conf *Config) logPath(id string) string {
	util.LogE(os.MkdirAll(Conf.LogDir, os.ModePerm)) //没有目录先建目录
	return filepath.Join(conf.LogDir, id) + ".log"   //文件名是任务的id
}

func (conf *Config) getCommand(id string) *exec.Cmd {
	return exec.Command("ping", "192.168.2.77", "-n", "10")
}
