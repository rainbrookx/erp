package handler

import "github.com/gin-gonic/gin"

type PlatformConfigHandler struct{}

func NewPlatformConfigHandler() *PlatformConfigHandler {
	return &PlatformConfigHandler{}
}

// GetPlatformName 获取平台名称
func (h *PlatformConfigHandler) GetPlatformName(c *gin.Context) {

}

// GetPlatformUrl 获取官方网站地址
func (h *PlatformConfigHandler) GetPlatformUrl(c *gin.Context) {

}

// GetPlatformRegisterFlag 获取是否开启注册
func (h *PlatformConfigHandler) GetPlatformRegisterFlag(c *gin.Context) {

}

// GetPlatformCheckCodeFlag 获取是否开启验证码
func (h *PlatformConfigHandler) GetPlatformCheckCodeFlag(c *gin.Context) {

}

// GetPlatformAppVersion 获取APP版本
func (h *PlatformConfigHandler) GetPlatformAppVersion(c *gin.Context) {

}
