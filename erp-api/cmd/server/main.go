package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/infrastructure/config"
	"github.com/rainbrookx/erp/internal/infrastructure/database"
	"github.com/rainbrookx/erp/internal/infrastructure/redis_infr"
	"github.com/rainbrookx/erp/internal/router"
	"github.com/rainbrookx/erp/internal/util"
)

func main() {
	initMain()

	engine := gin.Default()

	group := engine.Group("/jshERP-boot")
	router.InitRouter(group)

	err := engine.Run(":9999")
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func initMain() {
	config.InitConfig()

	database.InitDatabase(config.C.Database)
	redis_infr.InitRedis(config.C.Redis)
	util.InitRandomImageUtil(redis_infr.R)
}
