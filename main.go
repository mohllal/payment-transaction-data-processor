package main

import (
	router "github.com/Mohllal/go-challange/routers"

	"github.com/gin-gonic/gin"
)

var SetupRouter = func() *gin.Engine {
	r := gin.Default()
	router.PingRouter("api/", r)
	router.TransactionRouter("api/payment/", r)
	return r
}

func main() {
	r := SetupRouter()
	r.Run(":5000")
}
