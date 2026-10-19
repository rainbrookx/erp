package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/service"
)

type MaterialPropertyHandler struct{}

func NewMaterialPropertyHandler() *MaterialPropertyHandler {
	return &MaterialPropertyHandler{}
}

func (m *MaterialPropertyHandler) GetAllList(c *gin.Context) {

	list, err := service.ListMaterialProperty(c)
	if err != nil {
		c.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "获取数据失败",
		})
		return
	}

	c.JSON(200, dto.Success(list))
}
