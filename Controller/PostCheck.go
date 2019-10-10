package Controller

import (
	"../util"
	"github.com/kataras/iris"
)

func PostCheck(ctx iris.Context) {
	file, _, err := ctx.FormFile("jmx")
	if err != nil {
		util.LogE(err)
		responseMsg(ctx, iris.Map{"ok": false, "message": "请在Post的jmx字段上传您的jmx文件"})
		_ = file.Close()
	}
}
