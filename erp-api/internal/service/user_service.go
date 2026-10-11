package service

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"strings"
	"time"
	"uuid"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/repository"
	"github.com/rainbrookx/erp/internal/util"
)

func Login(ctx *gin.Context, loginName, password string) (*map[string]any, error) {
	data := make(map[string]any)
	msgTip := ""

	// 查询 session
	_, err := util.GetObjectFromSessionByKey(ctx, "userId")
	if err == nil {
		msgTip = "user already login"
	}

	// 删除旧的 session，之后换用新的 session
	util.DeleteObjectBySession(ctx, "userId")

	// 数据库查询，获取用户状态
	token := strings.ReplaceAll(uuid.New().String(), "-", "")

	msgTip, user, err := validateUser(ctx, loginName, password)
	if err != nil {
		return nil, err
	}

	if user.TenantID != nil {
		token = token + "_" + strconv.FormatInt(*user.TenantID, 10)
	}

	// 保存 session
	util.StorageObjectBySession(ctx, token, "userId", user.ID)

	// 判断密码是否为默认的
	pwdSimple := strings.EqualFold(util.MD5Crypto("123456"), util.MD5Crypto(*user.Password))
	// TODO 如果是管理员，则发送登录邮件

	// 组装数据返回
	data["msgTip"] = msgTip
	data["token"] = token
	data["user"] = user
	data["pwdSimple"] = pwdSimple
	return &data, nil
}

func validateUser(ctx context.Context, loginName, password string) (msgTip string, user *model.User, err error) {
	// 通过 loginName 查询是否存在用户（有逻辑删除）0未删除，1删除
	user, err = repository.GetUserByLoginNameWithDeleteFlag(ctx, loginName, "0")
	if err != nil {
		msgTip = err.Error()
		return msgTip, nil, errors.New(msgTip)
	}

	if *user.Status != 0 {
		msgTip = "user is black"
		return msgTip, nil, errors.New(msgTip)
	}

	// 查询租户信息
	tenant, err2 := repository.GetTenantByTenantId(ctx, *user.TenantID)
	if err2 != nil {
		msgTip = err2.Error()
		return msgTip, nil, errors.New(msgTip)
	}

	if tenant != nil {
		if !bytes.Equal(*tenant.Enabled, []byte{1}) {
			msgTip = "tenant is black"
			return msgTip, nil, errors.New(msgTip)
		}

		if time.Now().After(*tenant.ExpireTime) {
			msgTip = "tenant is expire"
			return msgTip, nil, errors.New(msgTip)
		}
	}

	// 判断账号密码（数据库存储的是前端对密码的MD5）
	if !strings.EqualFold(password, *user.Password) {
		msgTip = "user password error"
		return msgTip, nil, errors.New(msgTip)
	}

	msgTip = "user can login"
	return msgTip, user, nil
}
