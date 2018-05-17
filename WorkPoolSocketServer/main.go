package main

import (
	"github.com/kataras/iris"
	"fmt"
)

type Data struct {
	Task string
}

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./templates", ".html"))
	app.StaticWeb("/js", "./static/js") // serve our custom javascript code


	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	app.Post("/task", func(ctx iris.Context) {
		var task Data
		if err := ctx.ReadJSON(&task); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}
		fmt.Println(task)
		ctx.StatusCode(iris.StatusOK)

		//ctx.Writef("Received: %#+v\n", task)

	})
	app.Run(iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations)

}
