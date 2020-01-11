package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var GetPong = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
