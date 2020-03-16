package Task

import (
	"PressureMeter/util"
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type Config struct {
	JmxDir string `yaml:"JmxDir"` //存放jmx文件的目录位置
	JtlDir string `yaml:"JtlDir"` //存放jtl结果文件的目录位置
	LogDir string `yaml:"logDir"` //存放日志文件的目录位置
}

var Conf Config

var JmxDir = flag.String("JmxDir", "Data/jmx", "存放jmx文件的目录位置")
var JtlDir = flag.String("JtlDir", "Data/jtl", "存放jtl结果文件的目录位置")
var LogDir = flag.String("LogDir", "Data/log", "存放日志文件的目录位置")

func Init() {
	Conf.JmxDir = *JmxDir
	Conf.JtlDir = *JtlDir
	Conf.LogDir = *LogDir
	util.Log(".jmx file will be located in " + Conf.JmxDir)
	util.Log(".jtl file will be located in " + Conf.JtlDir)
	util.Log(".log file will be located in " + Conf.LogDir)
}

//通过id获取jmx文件路径
func (conf *Config) jmxPath(id string) string {
	util.LogE(os.MkdirAll(Conf.JmxDir, os.ModePerm)) //没有目录先建目录
	return filepath.Join(conf.JmxDir, id) + ".jmx"   //文件名是任务的id
}

//通过id获取jtl文件路径
func (conf *Config) jtlPath(id string) string {
	util.LogE(os.MkdirAll(Conf.JtlDir, os.ModePerm)) //没有目录先建目录
	return filepath.Join(conf.JtlDir, id) + ".jtl"   //文件名是任务的id
}

//通过id获取log文件路径
func (conf *Config) logPath(id string) string {
	util.LogE(os.MkdirAll(Conf.LogDir, os.ModePerm)) //没有目录先建目录
	return filepath.Join(conf.LogDir, id) + ".log"   //文件名是任务的id
}

//通过id获取要执行的指令
func (conf *Config) getCommand(id string) *exec.Cmd {
	return exec.Command("jmeter", "--nongui",
		"--testfile", conf.jmxPath(id),
		"--logfile", conf.jtlPath(id))
	/*
		return exec.Command("ping", "192.168.2.77", "-n", "10")
	*/
}

func getStopCommand(port int) *exec.Cmd {
	return exec.Command("shutdown.sh", strconv.Itoa(port))
}
