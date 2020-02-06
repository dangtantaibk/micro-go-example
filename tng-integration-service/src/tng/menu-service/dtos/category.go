package dtos

import "time"

type Category struct {
	ID               int64     `json:"category_id"`
	CategoryName     string    `json:"category_name"`
	CategoryValue    int64     `json:"category_value"`
	CategoryType     int64     `json:"category_type"`
	Order            int       `json:"order"`
	Status           int       `json:"status"`
	AreaID           int64     `json:"area_id"`
	ModifiedBy       string    `json:"modified_by"`
	ModifiedDateTime time.Time `json:"modified_date_time"`
}

type ListCategoryRequest struct {
	MerchantCode string `form:"merchant_code"`
}

type ListCategoryResponse struct {
	Meta Meta       `json:"meta"`
	Data []Category `json:"data"`
}

type CreateCategoryRequest struct {
	MerchantCode string `json:"merchant_code"`
	CategoryName string `json:"category_name"`
	Order        int    `json:"order"`
	Status       int    `json:"status"`
	AreaID       int64  `json:"area_id"`
}

type CreateCategoryResponse struct {
	Meta Meta `json:"meta"`
}

type UpdateStatusRequest struct {
	MerchantCode string `json:"merchant_code"`
	CategoryID   int64  `json:"category_id"`
	Status       int    `json:"status"`
}

type UpdateStatusResponse struct {
	Meta Meta `json:"meta"`
}

type DeleteCategoryRequest struct {
	CategoryID   int64 `json:"category_id"`
	MerchantCode string `json:"merchant_code"`
}

type DeleteCategoryResponse struct {
	Meta Meta `json:"meta"`
}

type UpdateCategoryRequest struct {
	MerchantCode string `json:"merchant_code"`
	CategoryName string `json:"category_name"`
	Order        int    `json:"order"`
	Status       int    `json:"status"`
	AreaID       int64  `json:"area_id"`
	CategoryID   int64  `json:"category_id"`
}

type UpdateCategoryResponse struct {
	Meta Meta `json:"meta"`
}
