// //go:build:wireinjection

package wire

// import (
// 	"github.com/google/wire"
// 	"pawtopia.com/internal/controller"
// 	"pawtopia.com/internal/repo"
// 	"pawtopia.com/internal/service"
// )

// func InitUserRouterHandler() (*controller.UserController, error) {
// 	wire.Build(
// 		repo.NewUserRepository,
// 		repo.NewUserAuthRepository,
// 		service.NewUserService,
// 		controller.NewUserController,
// 	)

// 	return new(controller.UserController), nil
// }
