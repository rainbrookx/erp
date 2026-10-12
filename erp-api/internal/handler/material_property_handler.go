package handler

import "github.com/gin-gonic/gin"

type MaterialPropertyHandler struct{}

func NewMaterialPropertyHandler() *MaterialPropertyHandler {
	return &MaterialPropertyHandler{}
}

func (m *MaterialPropertyHandler) GetAllList(c *gin.Context) {
	// TODO 未实现
}
