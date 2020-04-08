package main

import (
	"context"
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter"
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"github.com/kataras/iris"
	"time"
)

func main() {
	appCtx := context.Background()
	conf := PressureMeter.DefaultConfig()
	conf.LoggerConfig = &util.LoggerConfig{
		Logger: func(s string) { fmt.Printf("PressureMeter-->%s\n", s) },
		Error:  func(err error) { fmt.Printf("PressureMeter-->%s\n", err) },
	}
	conf.ModelConfig.DaemonConfig.UpdateStateCallback =
		func(runnings []Daemon.TaskInterface, startTimes []time.Time, queuings []Daemon.TaskInterface) {
			fmt.Println(runnings, startTimes, queuings)
		}
	app := PressureMeter.Init(appCtx, conf)
	err := app.Run(iris.Addr(":8080"))
	util.LogE(err)
}
