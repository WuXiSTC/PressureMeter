package PressureMeter

import (
	"context"
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/Model/Daemon"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"github.com/kataras/iris"
	"testing"
	"time"
)

func TestPressureMeter(t *testing.T) {
	ctx := context.Background()
	appCtx, _ := context.WithTimeout(ctx, 1000e9)
	conf := DefaultConfig()
	conf.LoggerConfig = &util.LoggerConfig{
		Logger: func(s string) { fmt.Printf("PressureMeter-->%s\n", s) },
		Error:  func(err error) { fmt.Printf("PressureMeter-->%s\n", err) },
	}
	conf.ModelConfig.DaemonConfig.UpdateStateCallback =
		func(runnings []Daemon.TaskInterface, startTimes []time.Time, queuings []Daemon.TaskInterface) {
			fmt.Println(runnings, startTimes, queuings)
		}
	app := Init(appCtx, conf)
	err := app.Run(iris.Addr(":80"))
	util.LogE(err)
}
