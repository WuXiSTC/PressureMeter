package Controller

import (
	"gitee.com/WuXiSTC/PressureMeter/Model"
	"gitee.com/WuXiSTC/PressureMeter/Model/TaskList"
	"github.com/kataras/iris"
	"time"
)

//启动一个任务
//
//返回启动是否成功和错误信息
func StartTask(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	duration, err := ctx.URLParamInt64("duration")
	if err != nil {
		responseMsg(ctx, iris.Map{"ok": false, "message": "duration格式错误"})
		return
	}
	if err := Model.TaskList.Start(taskId, time.Duration(duration)); err != nil {
		switch err.Error() {
		case "not exists":
			ctx.StatusCode(iris.StatusNotFound)
		case "already started":
			responseMsg(ctx, iris.Map{"ok": false, "message": "不能重复启动"})
		default:
			responseMsg(ctx, iris.Map{"ok": false, "message": err.Error()})
		}
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "启动成功"})
}

//停止一个任务
//
//返回停止是否成功和错误信息
func StopTask(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	errChan := make(chan error, 1)
	go func() {
		errChan <- Model.TaskList.Stop(taskId)
	}()
	select {
	case err := <-errChan:
		if err == nil {
			responseMsg(ctx, iris.Map{"ok": true, "message": "停止成功"})
			return
		}
		switch err.Error() {
		case "not exists":
			ctx.StatusCode(iris.StatusNotFound)
		default:
			responseMsg(ctx, iris.Map{"ok": false, "message": err.Error()})
		}
	case <-time.After(5e8): //如果停止任务耗时过长就尽快返回
		responseMsg(ctx, iris.Map{"ok": true, "message": "停止请求已提交"})
	}
}

func GetState(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	state := Model.TaskList.GetState(taskId)
	switch state {
	case TaskList.NOTEXISTS:
		ctx.StatusCode(iris.StatusNotFound)
	default:
		responseMsg(ctx, iris.Map{"message": Model.StateList[state], "stateCode": state})
	}
}
