package Controller

import (
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/Model"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"github.com/kataras/iris"
)

//新建任务，封装Service中的新建任务方法
func NewTask(ctx iris.Context) {
	file, _, err := ctx.FormFile("jmx") //字段检查，Post的jmx字段是否有文件上传
	if err != nil {
		util.LogE(err)
		responseMsg(ctx, iris.Map{"ok": false, "message": "请在Post的jmx字段上传您的jmx文件"})
		return
	}
	defer func() {
		_ = file.Close()
	}()

	taskId := ctx.Params().Get("id")
	if len(taskId) <= 0 {
		responseMsg(ctx, iris.Map{"ok": false, "message": "任务ID格式错误"})
		return
	}
	err = Model.AddNewTask(taskId, file)
	if err != nil {
		util.LogE(err)
		responseMsg(ctx, iris.Map{"ok": false, "message": err.Error()})
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "任务创建成功"})
}

//封装了Service中的删除任务函数Service.DeleteTask
//
//删除成功ok为true，否则ok为false并返回错误信息
func DeleteTask(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	if err := Model.TaskList.DelTask(taskId); err == nil {
		responseMsg(ctx, iris.Map{"ok": true, "message": "删除成功"})
		return
	} else if err.Error() == "not exists" {
		ctx.StatusCode(iris.StatusNotFound)
		return
	} else {
		responseMsg(ctx, iris.Map{"ok": false, "message": err.Error()})
	}
}

func ExpectDuration(ctx iris.Context) { //TODO
	_, err := ctx.WriteString(fmt.Sprintf("%d", 0))
	util.LogE(err)
}
