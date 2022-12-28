package controllers

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Ederene20/shortinx/helpers"
	"github.com/Ederene20/shortinx/validators"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func urls(c *gin.Context) {
	var urls []models.Url
	models.Database.Find(&urls)

	c.JSON(http.StatusOK, gin.H{"data": urls})
}

func shorten(c *gin.Context) {

	// Validate input
	var input validators.CreateUrlInput

	err := c.BindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	// Generate unique ID
	id, short_url := helpers.GenerateUniqueID()
	
	//
	url := models.Url{UniqueID: id, LongUrl: input.LongUrl, ShortUrl: short_url}
	models.Database.Create(&url)

	c.JSON(http.StatusOK, gin.H{"data": url})
}
