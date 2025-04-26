package routers

import (
	"pawtopia.com/internal/routers/admin"
	"pawtopia.com/internal/routers/user"
)

type RouterGroup struct {
	User  user.UserRouterGroup
	Admin admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
