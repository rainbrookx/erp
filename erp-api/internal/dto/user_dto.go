package dto

type LoginReq struct {
	LoginName string `json:"loginName,omitempty"`
	Password  string `json:"password,omitempty"`
	Uuid      string `json:"uuid,omitempty"`
	Code      string `json:"code,omitempty"`
}
