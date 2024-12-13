package config

import (
	"fmt"
	"go_server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitDB() {
	// 构建 DSN 字符串
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.Host,
		AppConfig.Database.Port,
		AppConfig.Database.Name,
		AppConfig.Database.Charset,
		AppConfig.Database.ParseTime,
		AppConfig.Database.Loc,
	)
	//驱动链接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库启动失败：%v", err)
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleCones)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenCones)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("数据库链接失败：%v", err)
	}

	global.DB = db
}
