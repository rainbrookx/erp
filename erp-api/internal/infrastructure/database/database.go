package database

import (
	"fmt"
	"log/slog"

	"github.com/rainbrookx/erp/internal/infrastructure/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(cfg config.DatabaseConfig) {
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		panic("连接数据库失败")
	}

	DB = db
}
