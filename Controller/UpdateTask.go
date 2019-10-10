package Controller

import (
	"../Service"
	"../util"
	"github.com/kataras/iris"
)

//新建任务
func UpdateTask(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	file, _, err := ctx.FormFile("jmx")
	if err != nil {
		return
	}
	defer util.LogE(file.Close())

	err = Service.UpdateTask(file, taskId)
	if err != nil {
		util.LogE(err)
		responseMsg(ctx, iris.Map{"ok": false, "message": err})
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "任务更新成功"})
}
