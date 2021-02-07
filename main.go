package main

import (
	"Iris-project/Controllers"
	"Iris-project/Crons"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/robfig/cron"
)

func main() {
	App := iris.New()
	mvc.Configure(App.Party("/user"), func(app *mvc.Application) {
		app.Handle(new(Controllers.UsersController))
	})
	App.Get("/", func(context *context.Context) {
		context.WriteString("Hello")
	})
	App.Get("/ping", func(context *context.Context) {
		//type myJson struct {
		//	Status int
		//	Msg string
		//}
		//json := myJson{200, "pong"}
		//context.JSON(json)
		context.JSON(iris.Map{"Msg": "pong"})
	})
	App.Run(iris.Addr(":8080"))

	c := cron.New()
	c.AddJob("*/3 * * * * ", &Crons.SendTask{})
	c.Start()
}
