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
	if exists := Model.TaskList.Start(taskId); !exists {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "启动成功"})
}

//停止一个任务
//
//返回停止是否成功和错误信息
func StopTask(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	if exists, err := Model.TaskList.Stop(taskId); !exists {
		ctx.StatusCode(iris.StatusNotFound)
		return
	} else if err != nil {
		responseMsg(ctx, iris.Map{"ok": false, "message": err.Error()})
		return
	} else {
		responseMsg(ctx, iris.Map{"ok": true, "message": "停止成功"})
	}
}

func GetState(ctx iris.Context) {
	taskId := ctx.Params().Get("id")
	if info, exists := Model.TaskList.GetInfo(taskId); !exists {
		ctx.StatusCode(iris.StatusNotFound)
		return
	} else {
		state := (*info).GetStateCode()
		responseMsg(ctx, iris.Map{"ok": true, "message": Model.StateList[state], "stateCode": state})
	}
}
