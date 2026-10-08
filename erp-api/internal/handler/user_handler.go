package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/util"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// RandomImage 获取随机校验码
func (h *UserHandler) RandomImage(c *gin.Context) {
	id, b64s, _, err := util.GenerateCaptcha()
	if err != nil {
		slog.Error(err.Error())
		c.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "获取失败",
		})
		return
	}

	c.JSON(200, dto.Success(gin.H{
		"base64": b64s,
		"uuid":   id,
	}))
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error(err.Error())
		c.String(500, "错误")
		return
	}

	c.JSON(200, dto.Success("登录成功"))
}
