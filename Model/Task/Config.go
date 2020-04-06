package Task

import (
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Config struct {
	JmxDir           string `yaml:"JmxDir" usage:"存放jmx文件的目录位置"`
	JtlDir           string `yaml:"JtlDir" usage:"存放jtl结果文件的目录位置"`
	LogDir           string `yaml:"logDir" usage:"存放日志文件的目录位置"`
	ShutdownBasePort uint16 `yaml:"BasePort" usage:"Jmeter用于接收shutdown message的端口号最大值"`

	taskAccN uint16
}

var conf Config

func DefaultConfig() Config {
	return Config{
		JmxDir:           "Data/jmx",
		JtlDir:           "Data/jtl",
		LogDir:           "Data/log",
		ShutdownBasePort: 4445,
		taskAccN:         4,
	}
}

func Init(c Config, cd Daemon.Config) {
	conf = c
	conf.taskAccN = cd.TaskAccN
}

//通过id获取jmx文件路径
func (conf *Config) jmxPath(id string) string {
	util.LogE(os.MkdirAll(conf.JmxDir, os.ModePerm)) //没有目录先建目录
	return filepath.Join(conf.JmxDir, id) + ".jmx"   //文件名是任务的id
}

//通过id获取jtl文件路径
func (conf *Config) jtlPath(id string) string {
	util.LogE(os.MkdirAll(conf.JtlDir, os.ModePerm)) //没有目录先建目录
	return filepath.Join(conf.JtlDir, id) + ".jtl"   //文件名是任务的id
}

//通过id获取log文件路径
func (conf *Config) logPath(id string) string {
	util.LogE(os.MkdirAll(conf.LogDir, os.ModePerm)) //没有目录先建目录
	return filepath.Join(conf.LogDir, id) + ".log"   //文件名是任务的id
}

func (conf *Config) startCommand(id string, i uint16) *exec.Cmd {
	return conf.getStartCommand(id, conf.getShutdownPort(i), getHosts(i))
}

func (conf *Config) stopCommand(i uint16) *exec.Cmd {
	return conf.getStopCommand(conf.getShutdownPort(i))
}

func (conf *Config) getShutdownPort(i uint16) uint16 {
	return conf.ShutdownBasePort + i
}

//通过id获取要执行的指令
func (conf *Config) getStartCommand(id string, shutdownPort uint16, IPList []string) *exec.Cmd {
	if IPList != nil && len(IPList) > 0 {
		return exec.Command("jmeter", "--nongui",
			"--testfile", conf.jmxPath(id),
			"--logfile", conf.jtlPath(id),
			"--jmeterlogfile", conf.logPath(id),
			"--jmeterproperty", fmt.Sprintf("jmeterengine.nongui.port=%d", shutdownPort),
			"--jmeterproperty", fmt.Sprintf("jmeterengine.nongui.maxport=%d", shutdownPort),

			"--remotestart", strings.Join(IPList, ","),
			"--jmeterproperty", "server.rmi.ssl.disable=true",
			"--remoteexit")
	} else {
		return exec.Command("jmeter", "--nongui",
			"--testfile", conf.jmxPath(id),
			"--logfile", conf.jtlPath(id),
			"--jmeterlogfile", conf.logPath(id),
			"--jmeterproperty", fmt.Sprintf("jmeterengine.nongui.port=%d", shutdownPort),
			"--jmeterproperty", fmt.Sprintf("jmeterengine.nongui.maxport=%d", shutdownPort))
	}
}

func (conf *Config) getStopCommand(shutdownPort uint16) *exec.Cmd {
	return exec.Command("shutdown.sh", fmt.Sprintf("%d", shutdownPort))
}
