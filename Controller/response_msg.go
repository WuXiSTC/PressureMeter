package Controller

import (
	"../util"
	"github.com/kataras/iris"
)

//统一的response函数
func responseMsg(ctx iris.Context, data iris.Map) {
	logResponse(ctx.JSON(data))
}

//记录response函数响应的结果
func logResponse(state int, err error) {
	if err != nil {
		util.LogE(err)
	}
}
