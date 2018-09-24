package main

import (
	"github.com/kataras/iris"
	"github.com/TechMaster/LearnGo/ChiREST/models"
)

func main() {
	app := iris.Default()
	app.Get("/articles", func(ctx iris.Context) {
		ctx.JSON(models.GetArticles())
	})
	app.Logger().SetLevel("error")
	app.Run(iris.Addr(":3333"))
}