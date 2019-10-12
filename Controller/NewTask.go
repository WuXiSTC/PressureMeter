package Controller

import (
	"../Service"
	"../util"
	"github.com/kataras/iris"
)

//新建任务
func NewTask(ctx iris.Context) {
	file, _, err := ctx.FormFile("jmx")
	if err != nil {
		util.LogE(err)
		responseMsg(ctx, iris.Map{"ok": false, "message": "请在Post的jmx字段上传您的jmx文件"})
		_ = file.Close()
		return
	}
	defer func() {
		_ = file.Close()
	}()

	taskId := ctx.Params().Get("id")
	err = Service.CreateTask(file, taskId)
	if err != nil {
		util.LogE(err)
		responseMsg(ctx, iris.Map{"ok": false, "message": err})
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "任务创建成功"})
}
