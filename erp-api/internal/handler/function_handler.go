package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/service"
)

type FunctionHandler struct{}

func NewFunctionHandler() *FunctionHandler {
	return &FunctionHandler{}
}

func (h *FunctionHandler) FindMenuByPNumber(c *gin.Context) {
	// TODO 未实现
	var reqMap map[string]any

	if err := c.ShouldBindJSON(&reqMap); err != nil {
		c.JSON(200, dto.Response[string]{
			Code:    500,
			Message: "JSON解析失败",
		})
		return
	}

	pNumber := reqMap["pNumber"].(string)
	userId := reqMap["userId"].(string)

	service.FindMenuByPNumber(c, pNumber, userId)
}
