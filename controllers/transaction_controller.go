package controllers

import (
	"code-challange/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

var allTransactions = data.LoadAllTransactions()

var GetTransactions = func(c *gin.Context) {
	// provider := c.Query("provider")
	// statusCode := c.Query("statusCode")
	// amountMin := c.Query("amountMin")
	// amountMax := c.Query("amountMax")
	// currency := c.Query("currency")

	// query := c.Request.URL.Query()
	// fmt.Println(query)

	// fmt.Println(flyATranscations.Transcations[0])
	// fmt.Println(flyBTranscations.Transcations[0])

	c.JSON(http.StatusOK, gin.H{
		"result": allTransactions,
	})
}
