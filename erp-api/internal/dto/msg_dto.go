package dto

import "github.com/rainbrookx/erp/internal/model"

// MsgSearch 对应 search 里面的json结构
type MsgSearch struct {
	Name string `json:"name"`
	// 其他筛选字段放这里
}

// MsgListQueryReq 外层query结构体
type MsgListQueryReq struct {
	SearchStr   string `form:"search"` // 先拿到原始json字符串
	Column      string `form:"column"`
	Order       string `form:"order"`
	FieldStr    string `form:"field"`
	CurrentPage int    `form:"currentPage"`
	PageSize    int    `form:"pageSize"`

	// 下面两个是解析后，业务用的，不参与query绑定
	Search MsgSearch `form:"-"`
	Field  []string  `form:"-"`
}

type MsgEx struct {
	model.Msg
	CreateTimeStr string `json:"createTimeStr"`
}
