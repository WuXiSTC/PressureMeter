package Controller

import (
	"../Model"
	"github.com/kataras/iris"
)

//启动一个任务
//
//返回启动是否成功和错误信息
func StartTask(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	err := Model.TaskList.Start(taskId)
	if err != nil {
		responseMsg(ctx, iris.Map{"ok": false, "message": err.Error()})
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "启动成功"})
}

//停止一个任务
//
//返回停止是否成功和错误信息
func StopTask(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	err := Model.TaskList.Stop(taskId)
	if err != nil {
		responseMsg(ctx, iris.Map{"ok": false, "message": err.Error()})
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "停止成功"})
}
