package dto

type MenuDto struct {
	Id        *int64    `json:"id"`
	Text      *string   `json:"text"`
	Icon      *string   `json:"icon"`
	Url       *string   `json:"url"`
	Component *string   `json:"component"`
	Children  []MenuDto `json:"children,omitempty"`
}

type FindMenuReq struct {
	PNumber int64 `json:"pNumber"`
	UserId  int64 `json:"userId"`
}
