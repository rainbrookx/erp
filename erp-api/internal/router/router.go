package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/handler"
)

func InitRouter(router *gin.RouterGroup) {
	bindUserHandler(router)
	bindPlatformConfigHandler(router)
	bindFunctionHandler(router)
	bindMaterialPropertyHandler(router)
	bindSystemConfigHandler(router)
	bindMsgHandler(router)
	bindDepotHeadHandler(router)
	bindDepotItemHandler(router)
}

// bindUserHandler 用户管理
func bindUserHandler(router *gin.RouterGroup) {
	relativePath := "/user"

	r := router.Group(relativePath)
	h := handler.NewUserHandler()

	r.GET("/randomImage", h.RandomImage)
	r.POST("/login", h.Login)
	r.GET("/logout", h.Logout)
	r.GET("/infoWithTenant", h.InfoWithTenant)
	r.GET("/getUserBtnByCurrentUser", h.GetUserBtnByCurrentUser)
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
	r.GET("/getInfoByKey", h.GetInfoByKey)
}

// 菜单功能
func bindFunctionHandler(router *gin.RouterGroup) {
	relativePath := "/function"

	r := router.Group(relativePath)
	h := handler.NewFunctionHandler()

	r.POST("/findMenuByPNumber", h.FindMenuByPNumber)
}

func bindMaterialPropertyHandler(router *gin.RouterGroup) {
	relativePath := "/materialProperty"

	r := router.Group(relativePath)
	h := handler.NewMaterialPropertyHandler()

	r.GET("/getAllList", h.GetAllList)
}

func bindSystemConfigHandler(router *gin.RouterGroup) {
	relativePath := "/systemConfig"

	r := router.Group(relativePath)
	h := handler.NewSystemConfigHandler()

	r.GET("/getCurrentInfo", h.GetCurrentInfo)
}

func bindMsgHandler(router *gin.RouterGroup) {
	relativePath := "/msg"

	r := router.Group(relativePath)
	h := handler.NewMsgHandler()

	r.GET("/list", h.List)
	r.GET("/getMsgByStatus", h.GetMsgByStatus)
}

func bindDepotHeadHandler(router *gin.RouterGroup) {
	relativePath := "/depotHead"

	r := router.Group(relativePath)
	h := handler.NewDepotHeadHandler()

	r.GET("/getBuyAndSaleStatistics", h.GetBuyAndSaleStatistics)
}

func bindDepotItemHandler(router *gin.RouterGroup) {
	relativePath := "/depotItem"

	r := router.Group(relativePath)
	h := handler.NewDepotItemHandler()

	r.GET("/buyOrSalePrice", h.BuyOrSalePrice)
}
