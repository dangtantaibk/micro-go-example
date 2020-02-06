package dtos

type ItemType struct {
	ID int64 `json:"item_type_id"`
	ItemTypeName string `json:"item_type_name"`
	Status int `json:"status"`
}

type ListItemTypeRequest struct {
	MerchantCode string `form:"merchant_code"`
}

type ListItemTypeResponse struct {
	Meta Meta `json:"meta"`
	Data []ItemType `json:"data"`
}

type CreateItemTypeRequest struct {
	MerchantCode string `json:"merchant_code"`
	ItemTypeName string `json:"item_type_name"`
	Status int `json:"status"`
}

type CreateItemTypeResponse struct {
	Meta Meta `json:"meta"`
	Data ItemType `json:"data"`
}

type UpdateItemTypeRequest struct {
	MerchantCode string `json:"merchant_code"`
	ItemTypeID int64 `json:"item_type_id"`
	ItemTypeName string `json:"item_type_name"`
	Status int `json:"status"`
}

type UpdateItemTypeResponse struct {
	Meta Meta `json:"meta"`
	Data ItemType `json:"data"`
}

type DeleteItemTypeRequest struct {
	MerchantCode string `json:"merchant_code"`
	ItemTypeID   int64 `json:"item_type_id"`
}

type DeleteItemTypeResponse struct {
	Meta Meta `json:"meta"`
}