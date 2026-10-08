package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/handler"
)

func InitRouter(router *gin.Engine) {
	bindUserGroup(router)
}

func bindUserGroup(router *gin.Engine) {
	relativePath := "/user"

	r := router.Group(relativePath)
	h := handler.NewUserHandler()

	r.GET("/randomImage", h.RandomImage)
	r.POST("/login", h.Login)
}
