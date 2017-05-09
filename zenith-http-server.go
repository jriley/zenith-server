package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"github.com/jriley/zenith-server/handlers"
)

func hello(context *iris.Context) {
	context.Writef("hello from %s", context.Path())
}

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	app.OnError(iris.StatusNotFound, func(context *iris.Context) {
		context.HTML(iris.StatusNotFound, "<h1>Custom not found handler </h>")
	})

	app.Get("/", hello)

	app.Get("/users/:userid", func(context *iris.Context) {
		context.Writef("hello user with id of %s", context.Param("userid"))
	})

	app.Get("myfiles/*file", func(context *iris.Context) {
		context.HTML(iris.StatusOK, "Hello from the dynamic path after /myfiles is: </br> <b>" +
			context.Param("file") +  "</b>")
	})

	app.Get("/users/:userid/messages/:messageid", func(context *iris.Context) {
		context.HTML(iris.StatusOK, `Message from the user with id:<br/> <b>` +
			context.Param("userid") +
			`</b> message id <b>` +
			context.Param("messageid") + `</b>`)
	})

	app.Post("/zenith", handlers.Foo)

	app.Listen(":8080")
}
