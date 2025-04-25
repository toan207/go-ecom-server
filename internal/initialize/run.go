package initialize

import (
	"fmt"

	"go.uber.org/zap"
	"pawtopia.com/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitMySQL()
	InitRedis()
	global.Logger.Info("Initialization complete", zap.String("ok", "success"))
	r := InitRouter()
	port := global.Config.Server.Port
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s...\n", port)
	r.Run(":" + port)
}
