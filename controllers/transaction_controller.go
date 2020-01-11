package controllers

import (
	"code-challange/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

var GetTransactions = func(c *gin.Context) {
	// load all transactions
	transactions := data.LoadAllTransactions()

	// get transactions by provider
	provider := c.Query("provider")
	if provider != "" {
		transactions = data.FindTransactionsByProvider(transactions, provider)
	}

	// get transactions by statusCode
	statusCode := c.Query("statusCode")
	if statusCode != "" {
		transactions = data.FindTransactionsByStatusCode(transactions, statusCode)
	}

	// amountMin := c.Query("amountMin")
	// amountMax := c.Query("amountMax")
	// currency := c.Query("currency")

	// query := c.Request.URL.Query()
	// fmt.Println(query)

	// fmt.Println(flyATranscations.Transcations[0])
	// fmt.Println(flyBTranscations.Transcations[0])

	c.JSON(http.StatusOK, gin.H{
		"result": transactions,
	})
}
