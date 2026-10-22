package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/model"
	"github.com/rainbrookx/erp/internal/query"
)

func SelectByConditionMsg(ctx *gin.Context, req dto.MsgListQueryReq, userInfo *model.User) ([]*model.Msg, error) {
	// 分页
	currentPage := req.CurrentPage
	pageSize := req.PageSize

	m := query.Q.Msg

	return m.WithContext(ctx).
		Where(
			m.DeleteFlag.Neq("1"),
			m.UserID.Eq(userInfo.ID),
			m.TenantID.Eq(*userInfo.TenantID),
			m.MsgTitle.Like("%"+req.Search.Name+"%"),
		).
		Order(m.CreateTime.Desc()).
		Limit(pageSize).
		Offset((currentPage - 1) * pageSize).
		Find()
}
