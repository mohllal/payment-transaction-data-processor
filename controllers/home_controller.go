package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var GetHome = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
