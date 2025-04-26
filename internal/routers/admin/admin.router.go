package admin

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	adminRouterPublic := Router.Group("admin")
	{
		adminRouterPublic.GET("/login")
	}

	adminRouterPrivate := Router.Group("admin/user")
	{
		adminRouterPrivate.POST("/create")
		adminRouterPrivate.POST("/activate")
	}
}
