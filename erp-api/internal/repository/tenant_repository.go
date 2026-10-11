package repository

import (
	"context"
	"errors"

	"github.com/rainbrookx/erp/internal/infrastructure/database"
	"github.com/rainbrookx/erp/internal/model"
)

func GetTenantByTenantId(ctx context.Context, tenantID int64) (*model.Tenant, error) {
	var tenant model.Tenant

	if err := database.DB.
		WithContext(ctx).
		Where(map[string]any{
			"tenant_id": tenantID,
		}).First(&tenant).Error; err != nil {
		return nil, errors.New("access service error")
	}

	return &tenant, nil
}
