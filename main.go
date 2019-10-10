package main

import (
	"./Controller"
	"./util"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

//main 主函数
func main() {

	app := iris.New()
	app.Use(logger.New())
	app.Post("/Task/new", Controller.PostCheck)
	app.Post("/Task/update", Controller.PostCheck)
	app.Post("/Task/new/{id:path}", Controller.NewTask)
	app.Post("/Task/update/{id:path}", Controller.UpdateTask)
	app.Get("/Task/delete/{id:path}", Controller.DeleteTask)
	//app.Get("/Task/start/{id:path}", Controller.RunJMX)
	//app.Get("/Task/stop/{id:path}", Controller.StopJMX)
	//app.Get("/Task/getConfig/{id:path}", Controller.GetJMX)
	//app.Get("/Task/getState/{id:path}", Controller.GetState)
	//app.Get("/Task/getResult/{id:path}", Controller.GetJTL)

	err := app.Run(iris.Addr(":8080"))
	util.LogE(err)
}
