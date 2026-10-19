package handler

import (
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/service"
)

type FunctionHandler struct{}

func NewFunctionHandler() *FunctionHandler {
	return &FunctionHandler{}
}

func (h *FunctionHandler) FindMenuByPNumber(c *gin.Context) {
	var req dto.FindMenuReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, dto.Response[string]{
			Code:    500,
			Message: "JSON解析失败",
		})
		return
	}

	pNumber := strconv.FormatInt(req.PNumber, 10)
	userId := strconv.FormatInt(req.UserId, 10)

	dataArray, err := service.FindMenuByPNumber(c, pNumber, userId)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "获取失败",
		})
		return
	}

	c.JSON(200, dto.Success(dataArray))
}
