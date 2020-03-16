package PressureMeter

import (
	"PressureMeter/util"
	"context"
	"fmt"
	"github.com/kataras/iris"
	"testing"
)

func TestPressureMeter(t *testing.T) {
	ctx := context.Background()
	appCtx, _ := context.WithTimeout(ctx, 20e9)
	conf := DefaultConfig()
	conf.LoggerConfig = &util.LoggerConfig{
		Logger: func(s string) { fmt.Printf("PressureMeter-->%s\n", s) },
		Error:  func(err error) { fmt.Printf("PressureMeter-->%s\n", err) },
	}
	app := Init(appCtx, conf)
	err := app.Run(iris.Addr(":8080"))
	util.LogE(err)
}
