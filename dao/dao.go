package dao

import (
	"gin-ranking/config"
	"gin-ranking/pkg/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	 Db *gorm.DB
	 err error
)

func init() {
	dsn := config.Mysqldb
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error!!!!": err}, "数据库连接失败")
	}

	sqlDB, err := Db.DB()
	if err != nil {
		logger.Error(map[string]interface{}{"datebase error!!!!": Db.Error}, "数据库连接失败")
	}
		// 连接池相关信息
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}