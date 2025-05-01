package user

import (
	"github.com/gin-gonic/gin"
	"pawtopia.com/internal/wire"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := Router.Group("user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.GET("/login")
		userRouterPublic.GET("/otp")
	}

	userRouterPrivate := Router.Group("user")
	{
		userRouterPrivate.GET("/info")
		userRouterPrivate.GET("/update")
		userRouterPrivate.GET("/delete")
	}
}
