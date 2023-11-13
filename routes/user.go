package routers

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	var ctrl controllers.UserController

	route := router.Group("users")

	route.GET("/:id", ctrl.GetUser)
	route.POST("/", ctrl.Register)
}
