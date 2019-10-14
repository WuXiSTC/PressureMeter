package Controller

import (
	"../Model"
	"../util"
	"github.com/kataras/iris"
)

//新建任务，封装Service中的新建任务方法
func NewTask(ctx iris.Context) {
	file, _, err := ctx.FormFile("jmx") //字段检查，Post的jmx字段是否有文件上传
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
	err = Model.AddNewTask(taskId, file)
	if err != nil {
		util.LogE(err)
		responseMsg(ctx, iris.Map{"ok": false, "message": err.Error()})
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "任务创建成功"})
}
