package database

import (
	"fmt"
	"log"
	"os"
	"puzzle/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func InitMySQL() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Global.Mysql.Username,
		config.Global.Mysql.Password,
		config.Global.Mysql.Host,
		config.Global.Mysql.Port,
		config.Global.Mysql.Database,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 启用彩色打印
		},
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// SkipDefaultTransaction: true, // 禁用默认事务
	})

	if err != nil {
		fmt.Printf("数据库链接错误: %v", err)
	}

	if db.Error != nil {
		fmt.Printf("数据库链接错误: %v", db.Error)
	}

	// 设置数据库连接池参数
	sqlDB, _ := db.DB()
	// 设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(20)
}

func GetMySQL() *gorm.DB { return db }
