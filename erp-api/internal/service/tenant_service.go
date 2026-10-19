package service

import (
	"context"

	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/query"
)

func GetTenantByTenantId(ctx context.Context, tenantId int64) (*model.Tenant, error) {
	t := query.Q.Tenant

	return t.WithContext(ctx).
		Where(t.TenantID.Eq(tenantId), t.DeleteFlag.Neq("1")).
		First()
}
