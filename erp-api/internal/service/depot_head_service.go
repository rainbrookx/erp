package service

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/query"
	"github.com/rainbrookx/erp/internal/util"
	"github.com/rainbrookx/erp/internal/util/tools"
	"gorm.io/gen/field"
)

type DepotHeadService struct {
}

var depotHeadDto *dto.DepotHeadDto

func NewDepotHeadService() *DepotHeadService {
	depotHeadDto = dto.NewDepotHeadDto()
	return &DepotHeadService{}
}

func (s *DepotHeadService) GetBuyAndSaleStatistics(ctx *gin.Context, today, monthFirstDay, yesterdayBegin, yesterdayEnd, yearBegin, yearEnd string) (map[string]any, error) {
	user, err := util.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	role, err := GetRoleTypeByUserId(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	var priceLimit string
	if role != nil && role.PriceLimit != nil {
		priceLimit = *role.PriceLimit
	}

	config, err := GetSystemConfig(ctx)
	if err != nil {
		return nil, err
	}
	var forceFlag = *config.ForceApprovalFlag == "1"
	var creatorArray = GetCreatorArray(ctx)

	var inOutPriceVoList = GetBuyAndSaleStatisticsList(ctx, yearBegin, yearEnd, creatorArray, forceFlag)

	periods := []string{"today", "month", "yesterday", "year"}
	types := []string{"Buy", "BuyBack", "Sale", "SaleBack", "RetailSale", "RetailSaleBack"}

	statistics := make(map[string]float64)

	for _, period := range periods {
		for _, type_ := range types {
			statistics[period+type_] = 0
		}
	}

	todayDate, _ := tools.StrToDate(today)
	monthFirstDate, _ := tools.StrToDate(monthFirstDay)
	yesterdayStartDate, _ := tools.StrToDate(yesterdayBegin)
	yesterdayEndDate, _ := tools.StrToDate(yesterdayEnd)
	yearStartDate, _ := tools.StrToDate(yearBegin)
	yearEndDate, _ := tools.StrToDate(yearEnd)

	for _, item := range inOutPriceVoList {
		operTime := item.OperTime
		discountLastMoney := item.DiscountLastMoney
		totalPriceAbs := math.Abs(item.TotalPrice)

		if isWithinRange(operTime, todayDate, time.Now()) {
			updateStatistics(statistics, item, "today", discountLastMoney, totalPriceAbs)
		}

		if isWithinRange(operTime, monthFirstDate, time.Now()) {
			updateStatistics(statistics, item, "month", discountLastMoney, totalPriceAbs)
		}

		if isWithinRange(operTime, yesterdayStartDate, yesterdayEndDate) {
			updateStatistics(statistics, item, "yesterday", discountLastMoney, totalPriceAbs)
		}

		if isWithinRange(operTime, yearStartDate, yearEndDate) {
			updateStatistics(statistics, item, "year", discountLastMoney, totalPriceAbs)
		}
	}

	result := make(map[string]any)

	for _, period := range periods {
		result[period+"Buy"] = parseHomePriceByLimit(
			statistics[period+"Buy"]-statistics[period+"BuyBack"],
			"buy",
			priceLimit,
			"***",
		)
		result[period+"Sale"] = parseHomePriceByLimit(
			statistics[period+"Sale"]-statistics[period+"SaleBack"],
			"sale",
			priceLimit,
			"***")
		result[period+"RetailSale"] = parseHomePriceByLimit(
			statistics[period+"RetailSale"]-statistics[period+"RetailSaleBack"],
			"retail",
			priceLimit,
			"***")
	}

	return result, nil
}

func GetCreatorArray(ctx *gin.Context) []string {
	creator := GetCreatorByCurrentUser(ctx)
	var creatorArray []string
	if creator != "" {
		creatorArray = strings.Split(creator, ",")
	}
	return creatorArray
}

func GetCreatorByCurrentUser(ctx *gin.Context) string {
	creator := ""
	user, err := util.GetCurrentUser(ctx)
	if err != nil {
		return creator
	}

	role, err := GetRoleTypeByUserId(ctx, user.ID)
	if err != nil || role == nil {
		return creator
	}

	roleType := *role.Type

	if "个人数据" == roleType {
		creator = strconv.FormatInt(user.ID, 10)
	} else if "本部门数据" == roleType {
		creator = GetUserIdListByUserId(ctx, user.ID)
	}

	return creator
}

func GetUserIdListByUserId(ctx *gin.Context, userId int64) string {
	users := ""
	ou := query.Q.OrgaUserRel

	our, err := ou.WithContext(ctx).
		Where(ou.UserID.Eq(userId)).First()
	if err != nil {
		return users
	}

	userIdList, err := GetUserIdListByOrgId(ctx, *our.OrgaID)
	if err != nil {
		return users
	}

	for _, u := range userIdList {
		users = users + strconv.FormatInt(u, 10) + ","
	}

	if "" != users {
		users = users[:len(users)-1]
	}
	return users
}

func GetUserIdListByOrgId(ctx *gin.Context, orgId int64) ([]int64, error) {
	orgIdList, err := GetOrgIdByParentId(ctx, orgId)
	if err != nil {
		return nil, err
	}

	var userIdList []int64

	ou := query.Q.OrgaUserRel

	ou.WithContext(ctx)
	if orgIdList != nil && len(orgIdList) > 0 {
		ou.Where(
			ou.OrgaID.In(orgIdList...),
			ou.DeleteFlag.Neq("1"),
		)
	} else {
		ou.Where(ou.DeleteFlag.Neq("1"))
	}

	list, err := ou.Find()
	if err != nil {
		return nil, err
	}

	if list != nil && len(list) > 0 {
		for _, our := range list {
			userIdList = append(userIdList, our.UserID)
		}
	}

	return userIdList, nil
}

func GetOrgIdByParentId(ctx *gin.Context, orgId int64) ([]int64, error) {
	var idList []int64

	o := query.Q.Organization

	orgList, err := o.WithContext(ctx).
		Where(o.DeleteFlag.Neq("1")).Find()
	if err != nil {
		return nil, err
	}

	if orgList != nil && len(orgList) > 0 {
		idList = append(idList, orgList[0].ID)
	}

	return idList, nil
}

func GetBuyAndSaleStatisticsList(ctx *gin.Context, yearBegin, yearEnd string, creatorArray []string, forceFlag bool) []dto.InOutPriceVo {
	dh := query.Q.DepotHead.As("dh")
	di := query.Q.DepotItem.As("di")

	// 构造查询
	queryExpr := dh.WithContext(ctx).
		Select(
			dh.ID,
			dh.DiscountLastMoney,
			dh.TotalPrice,
			dh.Type,
			dh.SubType,
			dh.OperTime,
		).
		Distinct().
		LeftJoin(di,
			// join on 条件：dh.id = di.header_id AND ifnull(di.delete_flag,'0') != '1'
			dh.ID.EqCol(di.HeaderID),
			field.Or(di.DeleteFlag.IsNull(), di.DeleteFlag.Neq("1")),
		).
		Where(
			// di.id is not null （把left join转成inner join效果）
			di.ID.IsNotNull(),
			// type in ('入库','出库')
			dh.Type.In("入库", "出库"),
			// ifnull(dh.delete_flag,'0') !='1'
			field.Or(dh.DeleteFlag.IsNull(), dh.DeleteFlag.Neq("1")),
		)

	// 动态条件
	if yearBegin != "" {
		t, _ := time.Parse("2006-01-02 15:04:05", yearBegin)
		queryExpr = queryExpr.Where(
			dh.OperTime.Gte(t))
	}

	if yearEnd != "" {
		t, _ := time.Parse("2006-01-02 15:04:05", yearEnd)
		queryExpr = queryExpr.Where(
			dh.OperTime.Lte(t))
	}

	if forceFlag {
		// dh.status in ('1','2','3')
		queryExpr = queryExpr.Where(dh.Status.In("1", "2", "3"))
	}

	if len(creatorArray) > 0 {
		queryExpr = queryExpr.Where(dh.Creator.In(util.StringsToInt64s(creatorArray)...))
	}

	// 执行查询，scan到vo
	var list []dto.InOutPriceVo
	err := queryExpr.Scan(&list)
	if err != nil {
		return list
	}
	return list
}

// isWithinRange 判断 operTime 是否在 [startDate, endDate] 闭区间内
func isWithinRange(operTime, startDate, endDate time.Time) bool {
	return !operTime.Before(startDate) && !operTime.After(endDate)
}

func updateStatistics(
	statistics map[string]float64,
	item dto.InOutPriceVo,
	period string,
	discountLastMoney float64,
	totalPriceAbs float64,
) {
	switch item.Type {
	case "入库":
		switch item.SubType {
		case "采购":
			statistics[period+"Buy"] += discountLastMoney
		case "销售退货":
			statistics[period+"SaleBack"] += discountLastMoney
		case "零售退货":
			statistics[period+"RetailSaleBack"] += totalPriceAbs
		}
	case "出库":
		switch item.SubType {
		case "采购退货":
			statistics[period+"BuyBack"] += discountLastMoney
		case "销售":
			statistics[period+"Sale"] += discountLastMoney
		case "零售":
			statistics[period+"RetailSale"] += totalPriceAbs
		}
	}
}

func parseHomePriceByLimit(price float64, typ, priceLimit, emptyInfo string) any {
	if priceLimit != "" {
		if typ == "buy" && strings.Contains(priceLimit, "1") {
			return emptyInfo
		}
		if typ == "retail" && strings.Contains(priceLimit, "2") {
			return emptyInfo
		}
		if typ == "sale" && strings.Contains(priceLimit, "3") {
			return emptyInfo
		}
	}
	return price
}
