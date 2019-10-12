package Controller

import (
	"../Service"
	"github.com/kataras/iris"
)

//封装了Service删除任务
func DeleteTask(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	err := Service.DeleteTask(taskId)
	if err != nil {
		responseMsg(ctx, iris.Map{"ok": false, "message": err})
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "删除成功"})
}
