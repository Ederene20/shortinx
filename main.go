package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Ederene20/shortinx/models"
	"github.com/Ederene20/shortinx/controllers"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	// Simple group: v1
	v1 := router.Group("/v1") {
		v1.GET("/ping", controllers.ping)
		v1.POST("/shorten", controllers.shorten)
		v1.GET("/urls", controllers.urls)
	}

	router.Run()
}
