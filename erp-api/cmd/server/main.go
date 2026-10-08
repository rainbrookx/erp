package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/router"
)

func main() {
	engine := gin.Default()

	group := engine.Group("/jshERP-boot")

	router.InitRouter(group)

	err := engine.Run(":9999")
	if err != nil {
		slog.Error(err.Error())
		return
	}
}
