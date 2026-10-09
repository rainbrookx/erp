package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/rainbrookx/erp/internal/infrastructure/database"
	"github.com/rainbrookx/erp/internal/model"
)

func GetPlatformConfigByKey(ctx context.Context, platformKey string) (*model.PlatformConfig, error) {
	if platformKey == "" {
		return nil, errors.New("不支持空字符串查询")
	}

	if strings.Contains(platformKey, "aliOss") ||
		strings.Contains(platformKey, "weixin") {
		return nil, errors.New("禁止查询：" + platformKey)
	}

	var m model.PlatformConfig

	if err := database.DB.
		WithContext(ctx).
		Where(model.PlatformConfig{PlatformKey: platformKey}).
		First(&m).
		Error; err != nil {
		return nil, err
	}

	return &m, nil
}
