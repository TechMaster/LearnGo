package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
	"flag"

	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
	"github.com/TechMaster/LearnGo/UploadConvertFeedback/encode_server"
	"github.com/TechMaster/LearnGo/UploadConvertFeedback/socket_server"
)

const maxSize = 50 << 20 // 5MB

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./templates", ".html"))
	app.StaticWeb("/js", "./static/js") // serve our custom javascript code

	socketserver := socket_server.New(app)
	var (
		maxQueueSize = flag.Int("queue_size", 100, "The size of job queue")
		maxWorkers   = flag.Int("workers", 3, "The number of workers to start")
		outputFolder = flag.String("output", "output", "Output folder")
	)
	flag.Parse()


	encode_server.New(*maxQueueSize, *maxWorkers, *outputFolder, socketserver)

	//----------Handle request from different paths ------
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	// Serve the upload_form.html to the client.
	app.Get("/upload", func(ctx iris.Context) {
		// create a token (optionally).
		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		// render the form with the token for any use you'd like.
		// ctx.ViewData("", token)
		// or add second argument to the `View` method.
		// Token will be passed as {{.}} in the template.
		ctx.View("upload_form.html", token)
	})



	// Handle the post request from the upload_form.html to the server
	app.Post("/upload", iris.LimitRequestBodySize(maxSize+1<<20), handleUpload)


	// Cấu hình CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.Use(cors)
	app.AllowMethods(iris.MethodOptions)

	// start the server at http://localhost:8080 with post limit at 5 MB.
	app.Run(iris.Addr(":8080") /* 0.*/, iris.WithPostMaxMemory(maxSize))
}
