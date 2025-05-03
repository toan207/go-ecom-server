package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gen"
	"pawtopia.com/global"
	"pawtopia.com/internal/po"
)

func checkErrorPanicC(err error, errorString string) {
	if err != nil {
		global.Logger.Error(errorString, zap.Error(err))
		panic(err)
	}
}

func InitMySQLC() {
	mysqlSetting := global.Config.MySQL
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, mysqlSetting.User, mysqlSetting.Password, mysqlSetting.Host, mysqlSetting.Port, mysqlSetting.Name)

	db, err := sql.Open("mysql", dsn)
	checkErrorPanicC(err, "Failed to connect to MySQL")
	global.Logger.Info("Connected to MySQL")
	global.MySQLC = db

	SetPoolC()
	global.Logger.Info("MySQL connection pool settings applied", zap.Int("MaxIdleConns", mysqlSetting.MaxIdleConns), zap.Int("MaxOpenConns", mysqlSetting.MaxOpenConns), zap.Int("ConnMaxLifetime", mysqlSetting.ConnMaxLifetime))
}

func SetPoolC() {
	mysqlSetting := global.Config.MySQL
	sqlDB := global.MySQLC

	if sqlDB == nil {
		checkErrorPanicC(nil, "Failed to get SQL DB from SQLC")
	}

	sqlDB.SetMaxIdleConns(mysqlSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlSetting.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlSetting.ConnMaxLifetime) * time.Second)
	global.Logger.Info("MySQL connection pool settings applied")

	// migrate()
	// genTableDAO()
}

func migrateC() {
	err := global.MySQL.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	checkErrorPanicC(err, "Failed to migrate database")

	global.Logger.Info("Database migration completed")
}

func genTableDAOC() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(global.MySQL)
	g.GenerateAllTable()
	g.Execute()
}
