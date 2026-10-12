package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/query"
)

func GetPlatformConfigByKey(ctx context.Context, platformKey string) (*model.PlatformConfig, error) {
	if platformKey == "" {
		return nil, errors.New("不支持空字符串查询")
	}

	if strings.Contains(platformKey, "aliOss") ||
		strings.Contains(platformKey, "weixin") {
		return nil, errors.New("禁止查询：" + platformKey)
	}

	var platformConfig *model.PlatformConfig

	c := query.Q.PlatformConfig

	platformConfig, err := c.WithContext(ctx).
		Where(c.PlatformKey.Eq(platformKey)).
		First()

	if err != nil {
		return nil, err
	}

	return platformConfig, nil
}
