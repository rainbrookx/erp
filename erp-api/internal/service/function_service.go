package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/repository"
	"github.com/rainbrookx/erp/internal/util"
)

func FindMenuByPNumber(ctx *gin.Context, pNumber, userId string) ([]dto.MenuDto, error) {
	var dataArray []dto.MenuDto
	var roleId int64 = 0
	fc := ""

	roleList, err := repository.ListUserBusinessByKeyIdAndType(ctx, userId, "UserRole")
	if err != nil {
		return nil, err
	}

	if roleList != nil && len(roleList) > 0 {
		value := *roleList[0].Value
		roleIdStr := strings.ReplaceAll(strings.ReplaceAll(value, "[", ""), "]", "")
		roleId, err = strconv.ParseInt(roleIdStr, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	// 当前用户所拥有的功能列表，格式如：[1][2][5]
	funList, err := repository.ListUserBusinessByKeyIdAndType(ctx, strconv.FormatInt(roleId, 10), "RoleFunctions")
	if err != nil {
		return nil, err
	}

	if funList != nil && len(funList) > 0 {
		fc = *funList[0].Value
	}

	// 获取系统配置信息-是否开启多级审核
	approvalFlag := "0"
	systemConfig, err := repository.GetSystemConfig(ctx)
	if err != nil {
		return nil, err
	}

	if systemConfig == nil {
		return nil, errors.New("systemConfig 为空")
	}

	approvalFlag = *systemConfig.MultiLevelApprovalFlag

	dataList, err := repository.ListFunctionByPNumber(ctx, pNumber)
	if err != nil {
		return nil, err
	}

	if dataList != nil && len(dataList) > 0 {
		userInfo, err2 := util.GetCurrentUser(ctx)
		if err2 != nil {
			return nil, err2
		}
		// 获取当前用户所属的租户所拥有的功能id的map
		funIdMap := make(map[int64]int64)

		funIdMap, err2 = util.GetCurrentTenantFunIdMap(ctx)
		if err2 != nil {
			return nil, err2
		}
		dataArray, err2 = repository.GetMenuByFunction(dataList, fc, approvalFlag, funIdMap, userInfo)
		if err2 != nil {
			return nil, err2
		}

		// 增加首页菜单项
		homeItem := dto.MenuDto{
			Id:        new(int64(0)),
			Text:      new("首页"),
			Icon:      new("home"),
			Url:       new("/dashboard/analysis"),
			Component: new("/layouts/TabLayout"),
		}
		dataArray = append([]dto.MenuDto{homeItem}, dataArray...)
	}

	return dataArray, nil
}
