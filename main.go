package main

import (
	router "code-challange/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.HomeRouter("api/", r)
	router.TransactionRouter("api/payment/transaction", r)
	r.Run(":5000")
}
