package handler

import (
	"log/slog"
	"time"

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

func (h *UserHandler) InfoWithTenant(c *gin.Context) {
	user, err := util.GetCurrentUser(c)
	if err != nil {
		c.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "获取失败",
		})
		return
	}

	// 获取当前用户数
	userCurrentNum, err := service.UserCount(c, *user.TenantID)
	if err != nil {
		c.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "获取失败",
		})
		return
	}

	tenant, err := service.GetTenantByTenantId(c, *user.TenantID)
	if err != nil {
		c.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "获取失败",
		})
		return
	}

	// 租户已经过期，移除token
	if tenant.ExpireTime != nil && time.Now().After(*tenant.ExpireTime) {
		util.DeleteObjectBySession(c, "userId")
		util.DeleteObjectBySession(c, "clientIp")
	}

	data := make(map[string]any)
	data["type"] = tenant.Type
	data["expireTime"] = util.ParseDateToStr(tenant.ExpireTime)
	data["userCurrentNum"] = userCurrentNum
	data["userNumLimit"] = tenant.UserNumLimit
	data["tenantId"] = tenant.TenantID
	c.JSON(200, dto.Success(data))
}

func (h *UserHandler) Logout(c *gin.Context) {
	util.DeleteObjectBySession(c, "userId")
	util.DeleteObjectBySession(c, "clientIp")

	c.JSON(200, dto.Response[any]{
		Code: 400,
	})
}

func (h *UserHandler) GetUserBtnByCurrentUser(c *gin.Context) {
	user, err := util.GetCurrentUser(c)

	if err != nil {
		c.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "获取失败",
		})
		return
	}

	userId := user.ID
	loginName := user.LoginName

	btnStrArr, err := service.GetBtnStrArrById(c, userId)

	if err != nil {
		c.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "获取失败",
		})
		return
	}

	data := make(map[string]any)
	if "admin" != loginName {
		data["userBtn"] = btnStrArr
	}

	c.JSON(200, dto.Success(data))
}
