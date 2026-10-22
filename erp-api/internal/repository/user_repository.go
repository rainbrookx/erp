package repository

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
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

func UserCount(ctx context.Context, tenantID int64) (int64, error) {
	u := query.Q.User

	return u.WithContext(ctx).
		Where(
			u.Status.Eq(0),
			u.DeleteFlag.Neq("1"),
			u.TenantID.Eq(tenantID),
		).
		Count()
}

func GetRoleWithoutTenant(ctx *gin.Context, roleId int64) (*model.Role, error) {
	r := query.Q.Role

	return r.WithContext(ctx).
		Where(
			r.DeleteFlag.Neq("1"),
			r.ID.Eq(roleId),
		).First()
}
