package Controller

import (
	"../Model"
	"github.com/kataras/iris"
)

//封装了Service中的删除任务函数Service.DeleteTask
//
//删除成功ok为true，否则ok为false并返回错误信息
func DeleteTask(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	err := Model.TaskList.DelTask(taskId)
	if err != nil {
		responseMsg(ctx, iris.Map{"ok": false, "message": err.Error()})
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "删除成功"})
}
