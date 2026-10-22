package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/service"
	"github.com/rainbrookx/erp/internal/util"
	"github.com/rainbrookx/erp/internal/util/tools"
)

type DepotHeadHandler struct {
}

var (
	depotHeadService *service.DepotHeadService
)

func NewDepotHeadHandler() *DepotHeadHandler {
	depotHeadService = service.NewDepotHeadService()

	return &DepotHeadHandler{}
}

func (h *DepotHeadHandler) GetBuyAndSaleStatistics(ctx *gin.Context) {
	user, err := util.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "错误",
		})
		return
	}

	if "admin" == user.LoginName {
		ctx.JSON(200, dto.Response[any]{
			Code: 200,
		})
		return
	}

	tool := &tools.Tools{}

	today := tool.GetNow() + tools.DayFirstTime
	monthFirstDay := tool.FirstDayOfMonth(tool.GetCurrentMonth()) + tools.DayFirstTime
	yesterdayBegin := tool.GetYesterday() + tools.DayFirstTime
	yesterdayEnd := tool.GetYesterday() + tools.DayLastTime
	yearBegin := tool.GetYearBegin() + tools.DayFirstTime
	yearEnd := tool.GetYearEnd() + tools.DayLastTime

	mapRs, err := depotHeadService.GetBuyAndSaleStatistics(ctx, today, monthFirstDay, yesterdayBegin, yesterdayEnd, yearBegin, yearEnd)
	ctx.JSON(200, dto.Response[any]{
		Code: 200,
		Data: mapRs,
	})
}
