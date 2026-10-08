package configs

import (
	"log/slog"

	"github.com/spf13/viper"
)

var C Config

// init 加载配置
func init() {
	v := viper.New()

	// 设置配置文件路径
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath("./configs")
	v.AddConfigPath(".")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		slog.Error(err.Error())
		panic("配置文件读取失败")
	}

	// 解析配置文件，反序列化到结构体
	if err := v.Unmarshal(&C); err != nil {
		slog.Error(err.Error())
		panic("配置解析文件失败")
	}

	// 校验配置文件
	if err := validate(&C); err != nil {
		slog.Error(err.Error())
		panic("配置文件校验失败")
	}
}

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	MySQL MySQLConfig `mapstructure:"mysql"`
}

// MySQLConfig MySQL 配置
type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

// RedisConfig Redis 配置
type RedisConfig struct {
}

func validate(c *Config) error {
	return nil
}
