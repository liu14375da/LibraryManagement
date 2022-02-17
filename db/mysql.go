package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitConn() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             3 * time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info,     // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,            // 禁用彩色打印
		},
	)

	conn, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		GlobalSettings.Datasource.Username,
		GlobalSettings.Datasource.Password,
		GlobalSettings.Datasource.Host,
		GlobalSettings.Datasource.Port,
		GlobalSettings.Datasource.Database)),
		&gorm.Config{
			// 禁用自动创建数据库外键约束
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   newLogger,
			DryRun:                                   false,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
		})
	//defer conn.Close()
	if err != nil {
		fmt.Println("gorm open err: ", err)
	}
	sqlDB, err := conn.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxIdleTime(time.Minute)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = conn
}
