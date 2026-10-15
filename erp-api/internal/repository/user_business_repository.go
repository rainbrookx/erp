package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/query"
)

func GetUserBusinessByKeyIdAndType(ctx *gin.Context, keyId, type_ string) ([]*model.UserBusiness, error) {
	u := query.Q.UserBusiness
	list, err := u.WithContext(ctx).
		Where(u.KeyID.Eq(keyId), u.Type.Eq(type_)).
		Find()

	if err != nil {
		return nil, err
	}

	return list, nil
}
