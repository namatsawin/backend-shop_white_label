package controllers

import (
	registerdto "backend/dto/user"
	"backend/services"

	gin "github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func (c UserController) GetUser(context *gin.Context) {
	response := c.userService.GetUserByID(context.Param("id"))
	context.JSON(response.Status, response)
}

func (c UserController) Register(context *gin.Context) {
	var form registerdto.RegisterDTO

	if validationErr := context.ShouldBindJSON(&form); validationErr != nil {
		response := new(registerdto.RegisterForm).Register(validationErr)
		context.JSON(response.Status, response)
		return
	}

	response := c.userService.Register(form)
	context.JSON(response.Status, response)
}
