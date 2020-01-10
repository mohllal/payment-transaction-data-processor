package routers

import (
	"code-challange/controllers"
	"github.com/gin-gonic/gin"
)

var TransactionRouter = func(baseUrl string, route *gin.Engine) {
	transcation := route.Group(baseUrl)
	{
		transcation.GET("/", controllers.GetTranscations)
	}
}
