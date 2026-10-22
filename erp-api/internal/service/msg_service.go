package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/repository"
	"github.com/rainbrookx/erp/internal/util"
)

func SelectMsg(ctx *gin.Context, req dto.MsgListQueryReq) ([]*dto.MsgEx, error) {
	var list []*dto.MsgEx

	userInfo, err := util.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	if "admin" != userInfo.LoginName {
		listMsg, err2 := repository.SelectByConditionMsg(ctx, req, userInfo)
		if err2 != nil {
			return nil, err2
		}

		if listMsg != nil {
			for _, msg := range listMsg {
				msgEx := dto.MsgEx{
					Msg:           *msg,
					CreateTimeStr: "",
				}

				if msg.CreateTime != nil {
					msgEx.CreateTimeStr = util.ParseDateToStr2(msgEx.CreateTime)
				}

				list = append(list, &msgEx)
			}
		}
	}

	return list, nil
}
