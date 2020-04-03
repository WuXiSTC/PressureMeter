package main

import (
	"context"
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"github.com/kataras/iris"
)

func main() {
	appCtx := context.Background()
	conf := PressureMeter.DefaultConfig()
	conf.LoggerConfig = &util.LoggerConfig{
		Logger: func(s string) { fmt.Printf("PressureMeter-->%s\n", s) },
		Error:  func(err error) { fmt.Printf("PressureMeter-->%s\n", err) },
	}
	app := PressureMeter.Init(appCtx, conf)
	err := app.Run(iris.Addr(":8080"))
	util.LogE(err)
}
