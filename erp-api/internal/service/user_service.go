package service

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
	"uuid"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
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
	pwdSimple := strings.EqualFold(util.MD5Crypto("123456"), *user.Password)
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

		trueBit := model.NullMyBitBool{
			MyBitBool: true,
			Valid:     true,
		}

		if *tenant.Enabled != trueBit {
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

func UserCount(ctx context.Context, tenantID int64) (int64, error) {
	return repository.UserCount(ctx, tenantID)
}

func GetBtnStrArrById(ctx *gin.Context, userId int64) ([]dto.BtnStrDto, error) {
	var btnStrArr []dto.BtnStrDto

	userRoleList, err := repository.ListUserBusinessByKeyIdAndType(ctx, strconv.FormatInt(userId, 10), "UserRole")
	if err != nil {
		return nil, err
	}

	if userRoleList != nil && len(userRoleList) > 0 {
		roleValue := *userRoleList[0].Value
		if roleValue != "" &&
			strings.Contains(roleValue, "[") &&
			strings.Contains(roleValue, "]") {

			roleValue = strings.ReplaceAll(strings.ReplaceAll(roleValue, "[", ""), "]", "")
			roleFunctionsList, err2 := repository.ListUserBusinessByKeyIdAndType(ctx, roleValue, "RoleFunctions")
			if err2 != nil {
				return nil, err2
			}

			if roleFunctionsList != nil && len(roleFunctionsList) > 0 {
				btnStr := *roleFunctionsList[0].BtnStr
				if btnStr != "" {
					err2 = json.Unmarshal([]byte(btnStr), &btnStrArr)
					if err2 != nil {
						return nil, err2
					}
				}
			}
		}
	}

	// 将数组中的funId转为url
	var btnStrWithUrlArr []dto.BtnStrDto
	if len(btnStrArr) > 0 {
		functionList, err2 := repository.ListAllFunction(ctx)
		if err2 != nil {
			return nil, err2
		}

		functionMap := make(map[int64]string)

		for _, function := range functionList {
			functionMap[function.ID] = *function.URL
		}

		for _, obj := range btnStrArr {
			funId := obj.FunId
			btnStr := obj.BtnStr

			btnStrWithUrlObj := dto.BtnStrDto{
				Url:    functionMap[funId],
				BtnStr: btnStr,
			}

			btnStrWithUrlArr = append(btnStrWithUrlArr, btnStrWithUrlObj)
		}
	}

	return btnStrWithUrlArr, nil
}
