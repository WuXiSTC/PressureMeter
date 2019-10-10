package Controller

import (
	"../Service"
	"../util"
	"github.com/kataras/iris"
)

//新建任务
func NewTask(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	file, _, err := ctx.FormFile("jmx")
	if err != nil {
		return
	}
	defer func() {
		util.LogE(file.Close())
	}()

	err = Service.CreateTask(file, taskId)
	if err != nil {
		util.LogE(err)
		responseMsg(ctx, iris.Map{"ok": false, "message": err})
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "任务创建成功"})
}
