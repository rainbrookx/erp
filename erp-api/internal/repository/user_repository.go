package repository

import (
	"context"
	"errors"

	"github.com/rainbrookx/erp/internal/infrastructure/database"
	"github.com/rainbrookx/erp/internal/model"
	"gorm.io/gorm"
)

// GetUserByLoginNameWithDeleteFlag 通过 loginName 查询是否存在用户（有逻辑删除）0未删除，1删除
func GetUserByLoginNameWithDeleteFlag(ctx context.Context, loginName, deleteFlag string) (*model.User, error) {
	if loginName == "" || deleteFlag == "" {
		return nil, errors.New("access service error")
	}

	var user model.User

	err := database.DB.
		WithContext(ctx).
		Where(map[string]any{
			"login_name":  loginName,
			"delete_flag": deleteFlag,
		}).
		First(&user).
		Error

	switch {
	case err == nil:
		return &user, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("user is not exist")
	default:
		return nil, errors.New("access service error")
	}
}
