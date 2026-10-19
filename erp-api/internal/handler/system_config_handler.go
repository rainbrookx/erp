package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/service"
)

type SystemConfigHandler struct{}

func NewSystemConfigHandler() *SystemConfigHandler {
	return &SystemConfigHandler{}
}

func (h SystemConfigHandler) GetCurrentInfo(c *gin.Context) {
	value, err := service.GetSystemConfig(c)
	if err != nil {
		c.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "获取数据失败",
		})
		return
	}

	c.JSON(200, dto.Success(value))
}
