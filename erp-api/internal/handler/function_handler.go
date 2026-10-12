package handler

import "github.com/gin-gonic/gin"

type FunctionHandler struct{}

func NewFunctionHandler() *FunctionHandler {
	return &FunctionHandler{}
}

func (h *FunctionHandler) FindMenuByPNumber(c *gin.Context) {
	// TODO 未实现
}
