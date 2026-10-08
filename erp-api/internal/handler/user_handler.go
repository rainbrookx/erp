package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/infrastructure/resp"
	"github.com/rainbrookx/erp/internal/infrastructure/tool"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) RandomImage(c *gin.Context) {
	id, b64s, _, err := tool.GenerateCaptcha()
	if err != nil {
		slog.Error(err.Error())
		c.String(500, "验证码生成失败")
		return
	}

	c.JSON(200, resp.Success(map[string]any{
		"base64": b64s,
		"uuid":   id,
	}))
}
