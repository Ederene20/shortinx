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
	router.POST(fmt.Sprint(api_version_1,"/shorten"), shorten)

	router.Run()
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func shorten(c *gin.Context) {
	var long_url Url

	err := c.BindJSON(&long_url)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, long_url)
}