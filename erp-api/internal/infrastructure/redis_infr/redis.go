package redis_infr

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/rainbrookx/erp/internal/infrastructure/config"
	"github.com/redis/go-redis/v9"
)

var R *redis.Client

func InitRedis(cfg config.RedisConfig) {
	opt := &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.Dbname,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	R = redis.NewClient(opt)

	// 连通性测试
	ctx := context.Background()
	res, err := R.Ping(ctx).Result()
	if err != nil {
		slog.Error(err.Error())
		panic("连接Redis失败")
	}

	slog.Info("连接Redis", "ping", res)
}

// Close 优雅关闭redis连接，程序退出时调用
func Close() error {
	if R != nil {
		return R.Close()
	}
	return nil
}
