package controller

import (
	"github.com/gin-gonic/gin"
	"pawtopia.com/internal/service"
	"pawtopia.com/pkg/response"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	email := c.PostForm("email")
	purpose := c.PostForm("purpose")
	result := uc.userService.Register(email, purpose)

	if result == response.ErrorCodeEmailExist {
		response.ErrorResponse(c, response.ErrorCodeEmailExist)
		return
	}

	response.SuccessResponse(c, response.ErrorCodeSuccess, nil)
}
