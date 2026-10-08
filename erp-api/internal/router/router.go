package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/handler"
)

func InitRouter(router *gin.RouterGroup) {
	bindUserHandler(router)
}

// bindUserHandler 用户管理
func bindUserHandler(router *gin.RouterGroup) {
	relativePath := "/user"

	r := router.Group(relativePath)
	h := handler.NewUserHandler()

	r.GET("/randomImage", h.RandomImage)
	r.POST("/login", h.Login)
}

// 平台参数
func bindPlatformConfigHandler(router *gin.RouterGroup) {
	relativePath := "/platformConfig"

	r := router.Group(relativePath)
	h := handler.NewPlatformConfigHandler()

	r.GET("/getPlatform/name", h.GetPlatformName)
	r.GET("/getPlatform/url", h.GetPlatformUrl)
	r.GET("/getPlatform/registerFlag", h.GetPlatformRegisterFlag)
	r.GET("/getPlatform/checkcodeFlag", h.GetPlatformCheckCodeFlag)
	r.GET("/getPlatform/appVersion", h.GetPlatformAppVersion)
}
