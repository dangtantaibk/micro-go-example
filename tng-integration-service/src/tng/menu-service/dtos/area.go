package dtos

type Area struct {
	ID int64 `json:"area_id"`
	AreaName string `json:"area_name"`
	Description string `json:"description"`
	Status int `json:"status"`
}

type ListAreaRequest struct {
	MerchantCode string `form:"merchant_code"`
}

type ListAreaResponse struct {
	Meta Meta `json:"meta"`
	Data []Area `json:"data"`
}

