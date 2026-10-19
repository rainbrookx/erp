package repository

import (
	"context"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/query"
)

func ListFunctionByPNumber(ctx *gin.Context, pNumber string) ([]*model.Function, error) {
	f := query.Q.Function

	enabledTrue := model.NullMyBitBool{
		MyBitBool: true,
		Valid:     true,
	}

	listFunction, err := f.WithContext(ctx).
		Where(f.Enabled.Eq(&enabledTrue),
			f.ParentNumber.Eq(pNumber),
			f.DeleteFlag.Neq("1")).
		Order(f.Sort).Find()

	if err != nil {
		return nil, err
	}

	return listFunction, nil
}

func GetFunctionFirstNodeIdList(ctx *gin.Context) ([]int64, error) {
	f := query.Q.Function

	enabledTrue := model.NullMyBitBool{
		MyBitBool: true,
		Valid:     true,
	}

	functions, err := f.WithContext(ctx).
		Select(f.ID).
		Where(
			f.ParentNumber.Eq("0"),
			f.Enabled.Eq(&enabledTrue),
			f.Or(
				f.DeleteFlag.IsNull(),
				f.DeleteFlag.Neq("1")),
		).
		Find()

	if err != nil {
		return nil, err
	}

	ids := make([]int64, 0, len(functions))
	for _, fn := range functions {
		ids = append(ids, fn.ID)
	}

	return ids, err
}

func GetMenuByFunction(dataList []*model.Function, fc, approvalFlag string, funIdMap map[int64]int64, userInfo *model.User) ([]dto.MenuDto, error) {
	var dataArray []dto.MenuDto

	for _, function := range dataList {
		// 如果不是超管也不是租户就需要校验，防止分配下级用户的功能权限，大于租户的权限
		_, hasFunctionId := funIdMap[function.ID]

		if "admin" == userInfo.LoginName ||
			userInfo.ID == *userInfo.TenantID ||
			hasFunctionId {
			// 如果关闭多级审核，遇到任务审核菜单直接跳过
			if "0" == approvalFlag && "/workflow" == *function.URL {
				continue
			}

			id := function.ID

			item := dto.MenuDto{
				Id:        &id,
				Text:      function.Name,
				Icon:      function.Icon,
				Url:       function.URL,
				Component: function.Component,
			}

			ctx := context.Background()
			newList, err := GetRoleFunction(&ctx, *function.Number)

			if err != nil {
				return nil, err
			}

			if len(newList) > 0 {
				childrenArr, err2 := GetMenuByFunction(newList, fc, approvalFlag, funIdMap, userInfo)

				if err2 != nil {
					return nil, err2
				}

				if len(childrenArr) > 0 {
					item.Children = childrenArr
					dataArray = append(dataArray, item)
				}

			} else {
				if strings.Contains(fc, "["+strconv.FormatInt(function.ID, 10)+"]") {
					dataArray = append(dataArray, item)
				}
			}
		}
	}
	return dataArray, nil

}

func GetRoleFunction(ctx *context.Context, pNumber string) ([]*model.Function, error) {
	f := query.Q.Function

	enabledTrue := model.NullMyBitBool{
		MyBitBool: true,
		Valid:     true,
	}

	return f.Where(
		f.Enabled.Eq(&enabledTrue),
		f.ParentNumber.Eq(pNumber),
		f.DeleteFlag.Neq("1"),
	).
		Order(f.Sort).
		Find()
}
