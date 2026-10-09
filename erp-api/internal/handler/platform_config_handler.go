package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/service"
)

type PlatformConfigHandler struct{}

func NewPlatformConfigHandler() *PlatformConfigHandler {
	return &PlatformConfigHandler{}
}

// GetPlatformName 获取平台名称
func (h *PlatformConfigHandler) GetPlatformName(c *gin.Context) {
	value, err := service.GetPlatformName(c)
	if err != nil {
		c.String(200, "ERP系统")
		return
	}

	c.String(200, value.PlatformValue)
}

// GetPlatformUrl 获取官方网站地址
func (h *PlatformConfigHandler) GetPlatformUrl(c *gin.Context) {
	value, err := service.GetPlatformUrl(c)

	if err != nil {
		c.String(200, "#")
		return
	}

	c.String(200, value.PlatformValue)
}

// GetPlatformRegisterFlag 获取是否开启注册
func (h *PlatformConfigHandler) GetPlatformRegisterFlag(c *gin.Context) {
	value, err := service.GetPlatformRegisterFlag(c)

	if err != nil {
		c.String(200, "#")
		return
	}

	c.String(200, value.PlatformValue)
}

// GetPlatformCheckCodeFlag 获取是否开启验证码
func (h *PlatformConfigHandler) GetPlatformCheckCodeFlag(c *gin.Context) {
	value, err := service.GetPlatformCheckCodeFlag(c)

	if err != nil {
		c.String(200, "#")
		return
	}

	c.String(200, value.PlatformValue)
}

// GetPlatformAppVersion 获取APP版本
func (h *PlatformConfigHandler) GetPlatformAppVersion(c *gin.Context) {
	value, err := service.GetPlatformAppVersion(c)

	if err != nil {
		c.String(200, "#")
		return
	}

	c.String(200, value.PlatformValue)
}
