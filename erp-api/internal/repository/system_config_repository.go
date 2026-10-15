package repository

import (
	"context"

	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/query"
)

func GetSystemConfig(ctx context.Context) (*model.SystemConfig, error) {
	s := query.Q.SystemConfig

	systemConfig, err := s.WithContext(ctx).
		Where(s.DeleteFlag.Neq("1")).
		First()

	if err != nil {
		return nil, err
	}

	return systemConfig, nil
}
