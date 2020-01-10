package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var GetHome = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
