package Controller

import (
	"../Model"
	"../Model/TaskList"
	"../util"
	"github.com/kataras/iris"
)

//获取配置文件并返回
func GetConfig(ctx iris.Context) {
	getFile(ctx, "log.log",
		func(taskInterface TaskList.TaskInfo) string {
			return taskInterface.GetConfigFilePath()
		})
}

//获取结果文件并返回
func GetResult(ctx iris.Context) {
	getFile(ctx, "log.log",
		func(taskInterface TaskList.TaskInfo) string {
			return taskInterface.GetResultFilePath()
		})
}

//获取日志文件并返回
func GetLog(ctx iris.Context) {
	getFile(ctx, "log.log",
		func(taskInterface TaskList.TaskInfo) string {
			return taskInterface.GetLogFilePath()
		})
}

//用函数式编程减少上面三个函数的代码量
func getFile(ctx iris.Context, filename string, getFileName func(taskInterface TaskList.TaskInfo) string) {
	taskId := ctx.Params().Get("id")
	task, exists := Model.TaskList.GetInfo(taskId)
	if !exists {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	path := getFileName(task)
	err := ctx.SendFile(path, filename)
	util.LogE(err)
}
