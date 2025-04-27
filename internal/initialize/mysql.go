package initialize

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"pawtopia.com/global"
	"pawtopia.com/internal/po"
)

func checkErrorPanic(err error, errorString string) {
	if err != nil {
		global.Logger.Error(errorString, zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	mysqlSetting := global.Config.MySQL
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, mysqlSetting.User, mysqlSetting.Password, mysqlSetting.Host, mysqlSetting.Port, mysqlSetting.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	checkErrorPanic(err, "Failed to connect to MySQL")
	global.Logger.Info("Connected to MySQL")
	global.MySQL = db

	SetPool()
	global.Logger.Info("MySQL connection pool settings applied", zap.Int("MaxIdleConns", mysqlSetting.MaxIdleConns), zap.Int("MaxOpenConns", mysqlSetting.MaxOpenConns), zap.Int("ConnMaxLifetime", mysqlSetting.ConnMaxLifetime))
}

func SetPool() {
	mysqlSetting := global.Config.MySQL
	sqlDB, err := global.MySQL.DB()

	checkErrorPanic(err, "Failed to get SQL DB from GORM")

	sqlDB.SetMaxIdleConns(mysqlSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlSetting.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlSetting.ConnMaxLifetime) * time.Second)
	global.Logger.Info("MySQL connection pool settings applied")

	// migrate()
	// genTableDAO()
}

func migrate() {
	err := global.MySQL.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	checkErrorPanic(err, "Failed to migrate database")

	global.Logger.Info("Database migration completed")
}

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(global.MySQL)
	g.GenerateAllTable()
	g.Execute()
}
