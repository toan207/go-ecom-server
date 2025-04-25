package initialize

import (
	"pawtopia.com/global"
	"pawtopia.com/pkg/logger"
)

func InitLogger() {
	logger := logger.NewLogger(global.Config.Logger)
	global.Logger = logger
}
