package routers

import (
	"github.com/Mohllal/go-challange/controllers"
	"github.com/gin-gonic/gin"
)

var TransactionRouter = func(baseUrl string, route *gin.Engine) {
	transaction := route.Group(baseUrl)
	{
		transaction.GET("/", controllers.GetTransactions)
	}
}
