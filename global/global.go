package global

import (
	"gorm.io/gorm"
	"pawtopia.com/pkg/logger"
	"pawtopia.com/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	MySQL  *gorm.DB
)
