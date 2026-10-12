package repository

import (
	"context"
	"errors"

	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/query"
)

func GetTenantByTenantId(ctx context.Context, tenantID int64) (*model.Tenant, error) {
	var tenant *model.Tenant

	t := query.Q.Tenant

	tenant, err := t.WithContext(ctx).
		Where(t.TenantID.Eq(tenantID)).
		First()

	if err != nil {
		return nil, errors.New("access service error")
	}

	return tenant, nil
}
