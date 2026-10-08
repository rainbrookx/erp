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

func (h *UserHandler) RandomImage(c *gin.Context) {
	id, b64s, _, err := util.GenerateCaptcha()
	if err != nil {
		slog.Error(err.Error())
		c.String(500, "验证码生成失败")
		return
	}

	c.JSON(200, dto.Success(map[string]any{
		"base64": b64s,
		"uuid":   id,
	}))
}

func (h *UserHandler) Login(context *gin.Context) {

}
