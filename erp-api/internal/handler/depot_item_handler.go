package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
)

type DepotItemHandler struct {
}

func NewDepotItemHandler() *DepotItemHandler {
	return &DepotItemHandler{}
}

func (h *DepotItemHandler) BuyOrSalePrice(ctx *gin.Context) {
	ctx.JSON(200, dto.Response[string]{
		Code: 500,
		Data: "接口展未实现",
	})
}
