package controllers

import (
	"code-challange/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var flyATranscations = data.LoadFlyATranscations()
var flyBTranscations = data.LoadFlyBTranscations()

var GetTranscations = func(c *gin.Context) {
	provider := c.Query("provider")

	fmt.Println(flyATranscations.Transcations[0])
	fmt.Println(flyBTranscations.Transcations[0])

	message := fmt.Sprintf("Hello from %s provider", provider)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
