package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/model"
)

func GetUserBusinessByKeyIdAndType(ctx *gin.Context, keyId, type_ string) ([]model.UserBusiness, error) {
	// todo 未实现
	return nil, nil
}
