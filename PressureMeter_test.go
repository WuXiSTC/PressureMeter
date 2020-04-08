package PressureMeter

import (
	"context"
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/Model/TaskList"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"github.com/kataras/iris"
	"testing"
)

func TestPressureMeter(t *testing.T) {
	ctx := context.Background()
	appCtx, _ := context.WithTimeout(ctx, 1000e9)
	conf := DefaultConfig()
	conf.LoggerConfig = &util.LoggerConfig{
		Logger: func(s string) { fmt.Printf("PressureMeter-->%s\n", s) },
		Error:  func(err error) { fmt.Printf("PressureMeter-->%s\n", err) },
	}
	conf.ModelConfig.UpdateStateCallback = func(list TaskList.TaskStateList) {
		fmt.Println(list)
	}
	app := Init(appCtx, conf)
	err := app.Run(iris.Addr(":80"))
	util.LogE(err)
}
