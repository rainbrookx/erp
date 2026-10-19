package repository

import (
	"context"

	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/query"
)

func ListMaterialProperty(ctx context.Context) ([]*model.MaterialProperty, error) {
	m := query.Q.MaterialProperty

	return m.WithContext(ctx).
		Where(m.DeleteFlag.Neq("1")).
		Find()
}
