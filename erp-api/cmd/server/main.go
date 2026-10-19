package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/infrastructure/config"
	"github.com/rainbrookx/erp/internal/infrastructure/database"
	"github.com/rainbrookx/erp/internal/infrastructure/redis_infr"
	"github.com/rainbrookx/erp/internal/query"
	"github.com/rainbrookx/erp/internal/router"
	"github.com/rainbrookx/erp/internal/util"
)

func main() {
	initMain()

	engine := gin.Default()

	group := engine.Group("/jshERP-boot")
	router.InitRouter(group)

	// 127.0.0.1 当前仅监听本机回环
	err := engine.Run("127.0.0.1:9998")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// TODO 可以考虑加上优雅退出
}

func initMain() {
	config.InitConfig()

	database.InitDatabase(config.C.Database)
	query.SetDefault(database.DB)
	redis_infr.InitRedis(config.C.Redis)
	util.InitRandomImageUtil(redis_infr.R)
	util.InitRedisUtil(redis_infr.R)
}
