package database

import (
	"fmt"
	"log/slog"

	"github.com/rainbrookx/erp/internal/infrastructure/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// todo 改写成 func init() 单例模式

func InitDatabase(cfg config.DatabaseConfig) (db *gorm.DB, err error) {
	mySQL := cfg.MySQL

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s",
		mySQL.Username,
		mySQL.Password,
		mySQL.Host,
		mySQL.Port,
		mySQL.Dbname,
		"utf8mb4",
		"Local",
	)
	slog.Info("连接数据库", "dsn", dsn)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
