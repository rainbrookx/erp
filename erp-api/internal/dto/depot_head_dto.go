package dto

import "time"

type DepotHeadDto struct {
}

func NewDepotHeadDto() *DepotHeadDto {
	return &DepotHeadDto{}
}

type InOutPriceVo struct {
	Id                int64     `json:"id"`
	DiscountLastMoney float64   `json:"discountLastMoney"`
	TotalPrice        float64   `json:"totalPrice"`
	Type              string    `json:"type"`
	SubType           string    `json:"subType"`
	OperTime          time.Time `json:"operTime"`
}
