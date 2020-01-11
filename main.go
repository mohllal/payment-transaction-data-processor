package main

import (
	router "code-challange/routers"

	"github.com/gin-gonic/gin"
)

var SetupRouters = func() *gin.Engine {
	r := gin.Default()
	router.HomeRouter("api/", r)
	router.TransactionRouter("api/payment/transaction", r)
	return r
}

func main() {
	r := SetupRouters()
	r.Run(":5000")
}
