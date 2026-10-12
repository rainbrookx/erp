package repository

import (
	"context"
	"errors"

	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/query"
	"gorm.io/gorm"
)

// GetUserByLoginNameWithDeleteFlag 通过 loginName 查询是否存在用户（有逻辑删除）0未删除，1删除
func GetUserByLoginNameWithDeleteFlag(ctx context.Context, loginName, deleteFlag string) (*model.User, error) {
	if loginName == "" || deleteFlag == "" {
		return nil, errors.New("access service error")
	}

	var user *model.User

	u := query.Q.User

	user, err := u.WithContext(ctx).
		Where(
			u.LoginName.Eq(loginName),
			u.DeleteFlag.Eq(deleteFlag),
		).
		First()

	switch {
	case err == nil:
		return user, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("user is not exist")
	default:
		return nil, errors.New("access service error")
	}
}
