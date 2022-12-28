package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Ederene20/shortinx/models"
	"github.com/Ederene20/shortinx/controllers"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	// Simple group: v1
	v1 := router.Group("api/v1") 
	{
		v1.GET("/ping", controllers.Ping)
		v1.GET("/:short_url", controllers.GetShortenedUrl)
		v1.POST("/shorten", controllers.Shorten)
		v1.GET("/urls", controllers.Urls)
	}

	router.Run()
}
