package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/router"
)

func main() {
	engine := gin.Default()

	router.InitRouter(engine)

	err := engine.Run(":8080")
	if err != nil {
		slog.Error(err.Error())
		return
	}
}
