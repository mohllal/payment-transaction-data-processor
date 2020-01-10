package controllers

import (
	"github.com/gin-gonic/gin"

	"fmt"

	"net/http"
)

var GetTranscations = func(c *gin.Context) {
	provider := c.Query("provider")

	message := fmt.Sprintf("Hello from %s provider", provider)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
