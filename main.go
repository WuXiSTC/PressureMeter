package main

import (
	"./Controller"
	"./Service"
	"./util"
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

//main 主函数
func main() {
	app := iris.New()
	iris.RegisterOnInterrupt(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 0)
		defer cancel()
		Service.Shutdown()
		util.LogE(app.Shutdown(ctx))
	})
	app.Use(logger.New())
	app.Post("/Task/new", Controller.PostCheck)
	app.Post("/Task/update", Controller.PostCheck)
	app.Post("/Task/new/{id:path}", Controller.NewTask)
	app.Post("/Task/update/{id:path}", Controller.UpdateTask)
	app.Get("/Task/delete/{id:path}", Controller.DeleteTask)
	app.Get("/Task/getConfig/{id:path}", Controller.GetConfig)
	app.Get("/Task/getResult/{id:path}", Controller.GetResult)
	app.Get("/Task/getLog/{id:path}", Controller.GetLog)
	app.Get("/Task/start/{id:path}", Controller.StartTask)
	//app.Get("/Task/stop/{id:path}", Controller.StopJMX)
	//app.Get("/Task/getState/{id:path}", Controller.GetState)

	err := app.Run(iris.Addr(":8080"))
	util.LogE(err)
}
