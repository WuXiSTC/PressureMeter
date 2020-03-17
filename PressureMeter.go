package PressureMeter

import (
	"context"
	"gitee.com/WuXiSTC/PressureMeter/Controller"
	"gitee.com/WuXiSTC/PressureMeter/Model"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"strings"
)

type Config struct {
	ModelConfig  Model.Config
	URLConfig    URLConfig
	LoggerConfig *util.LoggerConfig
}

func DefaultConfig() Config {
	return Config{
		ModelConfig: Model.DefaultConfig(),
		URLConfig: URLConfig{
			NewTask:    []string{"Task", "new"},
			DeleteTask: []string{"Task", "delete"},
			GetConfig:  []string{"Task", "getConfig"},
			GetResult:  []string{"Task", "getResult"},
			GetLog:     []string{"Task", "getLog"},
			StartTask:  []string{"Task", "start"},
			StopTask:   []string{"Task", "stop"},
			GetState:   []string{"Task", "getState"},
		},
	}
}

type URLConfig struct {
	NewTask    []string
	DeleteTask []string
	GetConfig  []string
	GetResult  []string
	GetLog     []string
	StartTask  []string
	StopTask   []string
	GetState   []string
}

func getURL(ss []string) string {
	return strings.Join(append(ss, "{id:path}"), "/")
}

//输入context和设置，返回一个iris app
func Init(ctx context.Context, conf Config) (app *iris.Application) {
	if conf.LoggerConfig != nil {
		util.SetLogger(*conf.LoggerConfig)
	}
	Model.Init(conf.ModelConfig)

	app = iris.New()
	app.Use(logger.New())
	app.Post(getURL(conf.URLConfig.NewTask), Controller.NewTask)
	app.Get(getURL(conf.URLConfig.DeleteTask), Controller.DeleteTask)
	app.Get(getURL(conf.URLConfig.GetConfig), Controller.GetConfig)
	app.Get(getURL(conf.URLConfig.GetResult), Controller.GetResult)
	app.Get(getURL(conf.URLConfig.GetLog), Controller.GetLog)
	app.Get(getURL(conf.URLConfig.StartTask), Controller.StartTask)
	app.Get(getURL(conf.URLConfig.StopTask), Controller.StopTask)
	app.Get(getURL(conf.URLConfig.GetState), Controller.GetState)

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(ctx, 10e9)
		defer cancel()
		util.LogE(app.Shutdown(ctx))
		util.LogE(Controller.Shutdown())
	}()
	return
}
