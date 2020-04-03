package Controller

import (
	"gitee.com/WuXiSTC/PressureMeter/Model"
	"gitee.com/WuXiSTC/PressureMeter/Model/TaskList"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"github.com/kataras/iris"
)

//获取配置文件并返回
func GetConfig(ctx iris.Context) {
	getFile(ctx, "jmx",
		func(taskInterface TaskList.TaskInfo) string {
			return taskInterface.GetConfigFilePath()
		})
}

//获取结果文件并返回
func GetResult(ctx iris.Context) {
	getFile(ctx, "jtl",
		func(taskInterface TaskList.TaskInfo) string {
			return taskInterface.GetResultFilePath()
		})
}

//获取日志文件并返回
func GetLog(ctx iris.Context) {
	getFile(ctx, "log",
		func(taskInterface TaskList.TaskInfo) string {
			return taskInterface.GetLogFilePath()
		})
}

//用函数式编程减少上面三个函数的代码量
func getFile(ctx iris.Context, suffix string, getFileName func(taskInterface TaskList.TaskInfo) string) {
	taskId := ctx.Params().Get("id")
	if task := Model.TaskList.GetInfo(taskId); task != nil {
		path := getFileName(task)
		err := ctx.SendFile(path, taskId+"."+suffix)
		util.LogE(err)
		return
	}
	ctx.StatusCode(iris.StatusNotFound)
}
