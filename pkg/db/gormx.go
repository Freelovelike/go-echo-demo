package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 开发环境开启 SQL 日志
	})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	println("数据库连接成功")

	// 设置连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间

	DB = db
}
