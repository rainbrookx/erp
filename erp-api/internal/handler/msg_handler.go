package handler

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/erp/internal/dto"
	"github.com/rainbrookx/erp/internal/service"
)

type MsgHandler struct{}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{}
}

func (h MsgHandler) List(c *gin.Context) {
	var req dto.MsgListQueryReq

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "请求参数错误",
		})
		return
	}

	// 手动解析 search 的json字符串到 Search 对象
	if req.SearchStr != "" {
		_ = json.Unmarshal([]byte(req.SearchStr), &req.Search)
	}

	// 手动分割 field 逗号字符串，并且过滤空元素
	if req.FieldStr != "" {
		parts := strings.SplitSeq(req.FieldStr, ",")
		for p := range parts {
			if p != "" {
				req.Field = append(req.Field, p)
			}
		}
	}

	msg, err := service.SelectMsg(c, req)
	if err != nil {
		c.JSON(500, dto.Response[string]{
			Code: 500,
			Data: "查询错误",
		})
		return
	}

	c.JSON(200, dto.Success(gin.H{
		"total": len(msg),
		"rows":  msg,
	}))
}
