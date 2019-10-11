package Controller

import (
	"../Service"
	"../util"
	"github.com/kataras/iris"
)

func GetConfig(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	path := Service.GetConfigFilePath(taskId)
	if path == "" {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	err := ctx.SendFile(path, "config.jmx")
	util.LogE(err)
}

func GetResult(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	path := Service.GetResultFilePath(taskId)
	if path == "" {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	err := ctx.SendFile(path, "result.jtl")
	util.LogE(err)
}

func GetLog(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	path := Service.GetLogFilePath(taskId)
	if path == "" {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	err := ctx.SendFile(path, "log.log")
	util.LogE(err)
}
