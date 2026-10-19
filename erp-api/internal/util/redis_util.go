package util

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/query"
	"github.com/rainbrookx/erp/internal/repository"
	"github.com/redis/go-redis/v9"
)

var bind = struct {
	client *redis.Client
}{}

const accessToken = "X-Access-Token"

// session的生命周期,秒
const maxSessionInSeconds = 60 * 60 * 24 * 3

func InitRedisUtil(rdb *redis.Client) {
	bind.client = rdb
}

func GetObjectFromSessionByKey(ctx *gin.Context, key string) (stringCmd *redis.StringCmd, err error) {
	if ctx == nil {
		return nil, errors.New("*gin.Context 为空")
	}

	token := ctx.GetHeader(accessToken)
	if token == "" {
		return nil, errors.New("X-Access-Token 为空")
	}

	result, err := bind.client.HExists(ctx, token, key).Result()
	if err != nil {
		return nil, err
	}

	if !result {
		return nil, errors.New("当前key不存在")
	}

	stringCmd = bind.client.HGet(ctx, token, key)
	bind.client.Expire(ctx, token, time.Duration(maxSessionInSeconds)*time.Second)

	return stringCmd, nil
}

func DeleteObjectBySession(ctx *gin.Context, key string) {
	if ctx != nil {
		token := ctx.GetHeader(accessToken)
		if token != "" {
			bind.client.HDel(ctx, token, key)
		}
	}
}

func StorageObjectBySession(ctx *gin.Context, token, key string, obj any) {
	bind.client.HSet(ctx, token, key, obj)
	bind.client.Expire(ctx, token, time.Duration(maxSessionInSeconds)*time.Second)
}

// GetCurrentUser 获取当前用户信息
func GetCurrentUser(ctx *gin.Context) (*model.User, error) {
	stringCmd, err := GetObjectFromSessionByKey(ctx, "userId")
	if err != nil {
		return nil, err
	}

	userId, err := stringCmd.Int64()
	if err != nil {
		return nil, err
	}

	u := query.Q.User
	user, err := u.WithContext(ctx).
		Where(u.ID.Eq(userId)).
		First()

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetCurrentTenantFunIdMap 获取当前用户所属的租户所拥有的功能id的map
func GetCurrentTenantFunIdMap(ctx *gin.Context) (map[int64]int64, error) {
	funIdMap := make(map[int64]int64)
	list, err := getCurrentTenantFunIdList(ctx)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		for _, funId := range list {
			funIdMap[funId] = funId
		}

		// 把一级菜单的id全都赋值给租户,解决漏掉枝干id的问题
		firstNodeIdList, err2 := repository.GetFunctionFirstNodeIdList(ctx)
		if err2 != nil {
			return nil, err2
		}

		for _, firstNodeId := range firstNodeIdList {
			funIdMap[firstNodeId] = firstNodeId
		}
		return funIdMap, nil
	}

	return nil, nil
}

// 获取当前用户所属的租户所拥有的功能id列表
func getCurrentTenantFunIdList(ctx *gin.Context) ([]int64, error) {
	funIdList := new([]int64)
	var roleId int64 = 0
	fc := ""
	userInfo, err := GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	// 获取当前用户所有的角色id
	roleList, err := repository.ListUserBusinessByKeyIdAndType(ctx, strconv.FormatInt(*userInfo.TenantID, 10), "UserRole")
	if err != nil {
		return nil, err
	}

	if roleList != nil && len(roleList) > 0 {
		value := *roleList[0].Value
		if value != "" {
			roleIdStr := strings.ReplaceAll(strings.ReplaceAll(value, "[", ""), "]", "")
			roleId, err = strconv.ParseInt(roleIdStr, 10, 64)
			if err != nil {
				return nil, err
			}
		}
	}

	// 当前用户所拥有的功能列表，格式如：[1][2][5]
	funList, err := repository.ListUserBusinessByKeyIdAndType(ctx, strconv.FormatInt(roleId, 10), "RoleFunctions")
	if err != nil {
		return nil, err
	}

	if funList != nil && len(funList) > 0 {
		fc = *funList[0].Value
	}

	if fc != "" {
		fc = fc[1 : len(fc)-1]
		fc = strings.ReplaceAll(fc, "][", ",")
		*funIdList = StrToLongList(fc)
	}

	return *funIdList, nil
}
