package controllers

import (
	"code-challange/data"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	// get transactions by amount
	amountMin := c.Query("amountMin")
	amountMax := c.Query("amountMax")
	if amountMin != "" && amountMax != "" {
		amountMin, _ := strconv.Atoi(amountMin)
		amountMax, _ := strconv.Atoi(amountMax)

		transactions = data.FindTransactionsByAmount(transactions, amountMin, amountMax)
	}

	// get transactions by currency
	currency := c.Query("currency")
	if currency != "" {
		transactions = data.FindTransactionsByCurrency(transactions, currency)
	}
	// query := c.Request.URL.Query()
	// fmt.Println(query)

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"result": transactions,
	})
}
