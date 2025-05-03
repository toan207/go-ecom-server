package global

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"pawtopia.com/pkg/logger"
	"pawtopia.com/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	MySQL  *gorm.DB
	Redis  *redis.Client
	MySQLC *sql.DB
)
