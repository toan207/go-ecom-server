package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"pawtopia.com/internal/service"
	"pawtopia.com/internal/vo"
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
	var params vo.UserRegistratorRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		response.ErrorResponse(c, response.ErrorCodeInvalidParams)
		return
	}

	result := uc.userService.Register(params.Email, params.Purpose)

	if result != response.ErrorCodeSuccess {
		response.ErrorResponse(c, result)
		return
	}

	response.SuccessResponse(c, response.ErrorCodeSuccess, nil)
}
