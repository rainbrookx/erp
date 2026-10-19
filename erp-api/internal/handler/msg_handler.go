package handler

import "github.com/gin-gonic/gin"

type MsgHandler struct{}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{}
}

func (h MsgHandler) List(c *gin.Context) {
	// TODO 未实现
}
