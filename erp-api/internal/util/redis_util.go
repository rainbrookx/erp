package util

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
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
