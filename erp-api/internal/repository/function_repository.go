package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/model"
)

func ListFunctionByPNumber(ctx *gin.Context, pNumber string) ([]model.Function, error) {
	//f := query.Q.Function
	//
	//f.WithContext(ctx).
	//	Where(f.Enabled.Eq())
	// todo 未实现
	return nil, nil
}
