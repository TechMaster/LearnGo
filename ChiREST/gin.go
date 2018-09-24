package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/TechMaster/LearnGo/ChiREST/models"
)

func setupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	gin.Logger()
	r := gin.New()

	r.GET("/articles", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.GetArticles())
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":3333")
}