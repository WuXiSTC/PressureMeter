package Controller

import (
	"../Service"
	"github.com/kataras/iris"
)

func StartTask(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	err := Service.StartTask(taskId)
	if err != nil {
		responseMsg(ctx, iris.Map{"ok": false, "message": err})
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "启动成功"})
}
