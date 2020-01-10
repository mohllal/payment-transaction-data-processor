package routers

import (
	"code-challange/controllers"
	"github.com/gin-gonic/gin"
)

var HomeRouter = func(baseUrl string, route *gin.Engine) {
	home := route.Group(baseUrl)
	{
		home.GET("/ping", controllers.GetHome)
	}
}
