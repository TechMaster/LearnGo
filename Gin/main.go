package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/TechMaster/LearnGo/Gin/models"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	//root path
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<h1>Hello World</h1>"))
	})

	//ping path
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "<h1>pong</h1>")
	})

	//pong returns pin
	r.GET("/pong", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html")
		c.String(http.StatusOK, "<h1>ping</h1>")
	})

	//cat returns a cat
	r.GET("/cat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "rock.html", nil)

	})

	// @Title getListArticles
	// @Description returns all articles
	// @Paths /articles
	// @Accept  json
	// @Param   customer_id     path    int     true        "Customer ID"
	// @Success 200 {array}  models.articles
	// @Failure 400 {object} "Articles not found"
	// @Resource model.articles
	// @Router /articles [get]

	r.GET("/articles", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.GetArticles())
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":3333")
}