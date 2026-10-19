package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/repository"
)

func GetSystemConfig(ctx *gin.Context) (*model.SystemConfig, error) {
	return repository.GetSystemConfig(ctx)
}
