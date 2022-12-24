package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Url struct {
	URL string `json:"url"`
}

func main() {
	router := gin.Default()
	api_version_1 := "api/v1"
	router.GET(fmt.Sprint(api_version_1,"/ping"), ping)

	router.Run()
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
