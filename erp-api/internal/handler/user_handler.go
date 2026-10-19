package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/service"
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
		c.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "用户登录失败",
		})
		return
	}

	// 校验验证码
	if !util.VerifyCaptcha(req.Uuid, req.Code) {
		c.JSON(200, dto.Response[any]{
			Code: 500,
			Data: gin.H{"message": "验证码错误"},
		})
		return
	}

	// 从数据库查询
	loginData, err := service.Login(c, req.LoginName, req.Password)
	if err != nil {
		c.JSON(200, dto.Response[any]{
			Code: 500,
			Data: gin.H{"message": err.Error()},
		})
		return
	}

	c.JSON(200, dto.Success(loginData))
}

func (h *UserHandler) InfoWithTenant(context *gin.Context) {
	// todo 未实现
}
