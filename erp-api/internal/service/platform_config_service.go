package service

import (
	"context"

	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/repository"
)

func GetPlatformName(ctx context.Context) (*model.PlatformConfig, error) {
	return repository.GetPlatformConfigByKey(ctx, "platform_name")
}

func GetPlatformUrl(ctx context.Context) (*model.PlatformConfig, error) {
	return repository.GetPlatformConfigByKey(ctx, "platform_url")
}

func GetPlatformRegisterFlag(ctx context.Context) (*model.PlatformConfig, error) {
	return repository.GetPlatformConfigByKey(ctx, "register_flag")

}

func GetPlatformCheckCodeFlag(ctx context.Context) (*model.PlatformConfig, error) {
	return repository.GetPlatformConfigByKey(ctx, "checkcode_flag")

}

func GetPlatformAppVersion(ctx context.Context) (*model.PlatformConfig, error) {
	return repository.GetPlatformConfigByKey(ctx, "app_version")
}
