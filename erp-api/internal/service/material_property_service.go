package service

import (
	"context"

	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/repository"
)

func ListMaterialProperty(ctx context.Context) ([]*model.MaterialProperty, error) {
	// 查询数据库：未删除的物料属性
	dbList, err := repository.ListMaterialProperty(ctx)
	if err != nil {
		return nil, err
	}

	// 构建内置预设属性列表（主数据，决定返回有哪些属性、顺序）
	var list []*model.MaterialProperty
	buildMaterialProperty(&list)

	mpMap := make(map[string]string)
	for _, mp := range dbList {
		if mp.NativeName != nil {
			mpMap[*mp.NativeName] = *mp.AnotherName
		}
	}

	// 遍历预设列表，匹配别名，覆盖
	for _, item := range list {
		if item.NativeName == nil {
			continue
		}
		if val, ok := mpMap[*item.NativeName]; ok {
			*item.AnotherName = val
		}
	}

	return list, nil
}

func buildMaterialProperty(list *[]*model.MaterialProperty) {
	mp1 := model.MaterialProperty{
		ID:          1,
		NativeName:  new("扩展1"),
		AnotherName: new("扩展1"),
	}

	mp2 := model.MaterialProperty{
		ID:          2,
		NativeName:  new("扩展2"),
		AnotherName: new("扩展2"),
	}

	mp3 := model.MaterialProperty{
		ID:          3,
		NativeName:  new("扩展3"),
		AnotherName: new("扩展3"),
	}

	*list = append(*list, &mp1, &mp2, &mp3)
}
