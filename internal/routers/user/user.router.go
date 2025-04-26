package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterPublic := Router.Group("user")
	{
		userRouterPublic.GET("/register")
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
