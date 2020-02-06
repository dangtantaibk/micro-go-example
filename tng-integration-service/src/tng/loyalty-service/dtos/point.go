package dtos

import (
	"tng/common/models/loyalty"
)

type PointInfo struct {
	Id int64 `json:"id"`

	UserId string `json:"user_id"`

	PointType string `json:"point_type"`

	Point int64 `json:"point"`

	Source string `json:"source"`

	ForTransactionId string `json:"for_transaction_id"`

	TransactionAmount int64 `json:"transaction_amount"`

	Notes string `json:"notes"`

	Created string `json:"created"`

	CreatedYmd string `json:"created_ymd"`

	Status int32 `json:"status"`

	AppId string `json:"app_id"`

	PromotionPercent float64 `json:"promotion_percent"`

	CampaignCode string `json:"campaign_code"`

	Channel string `json:"channel"`

	Rate float64 `json:"rate"`

	JsonDetail string `json:"json_detail"`
}

type InsertOrUpdatePointRequest struct {
	Id int64 `json:"id"`

	UserId string `json:"user_id"`

	PointType string `json:"point_type"`

	Point int64 `json:"point"`

	Source string `json:"source"`

	ForTransactionId string `json:"for_transaction_id"`

	TransactionAmount int64 `json:"transaction_amount"`

	Notes string `json:"notes"`

	Created string `json:"created"`

	CreatedYmd string `json:"created_ymd"`

	Status int32 `json:"status"`

	AppId string `json:"app_id"`

	PromotionPercent float64 `json:"promotion_percent"`

	CampaignCode string `json:"campaign_code"`

	Channel string `json:"channel"`

	Rate float64 `json:"rate"`

	JsonDetail string `json:"json_detail"`
}

type InsertOrUpdatePointResponse struct {
	Meta Meta `json:"meta"`
}

type DeletePointRequest struct {
	Id int64 `json:"id"`
}

type DeletePointResponse struct {
	Meta Meta `json:"meta"`
}

type ListPointRequest struct {
	PageIndex int32 `form:"page_index"`
	PageSize  int32 `form:"page_size"`
}

type ListPointResponse struct {
	Meta Meta         `json:"meta"`
	Data []*PointInfo `json:"data"`
}

type GetPointRequest struct {
	Id int64 `form:"id"`
}

type GetPointResponse struct {
	Meta Meta       `json:"meta"`
	Data *PointInfo `json:"data"`
}

type AddPointRequest struct {
	UserId            string  `json:"user_id"`
	Source            string  `json:"source"`
	ForTransactionId  string  `json:"for_transaction_id"`
	TransactionAmount int64   `json:"transaction_amount"`
	Notes             string  `json:"notes"`
	AppId             string  `json:"app_id"`
	PromotionPercent  float64 `json:"promotion_percent"`
	CampaignCode      string  `json:"campaign_code"`
	Channel           string  `json:"channel"`
	JsonDetail        string  `json:"json_detail"`
}
type AddPointResponse struct {
	Meta Meta        `json:"meta"`
	Data *WalletInfo `json:"data"`
}
type CheckPointRequest struct {
	UserId string `json:"user_id"`
}
type CheckPointResponse struct {
	Meta Meta        `json:"meta"`
	Data *WalletInfo `json:"data"`
}

type PointHistoryRequest struct {
	PageIndex int32  `json:"page_index,omitempty"`
	PageSize  int32  `json:"page_size,omitempty"`
	UserId    string `json:"user_id"`
}

type PointHistoryResponse struct {
	Meta Meta         `json:"meta"`
	Data []*PointInfo `json:"data"`
}

type CheckOldPointRequest struct {
	UserId           string `json:"user_id"`
	ForTransactionId string `json:"for_transaction_id"`
	AppId            string `json:"app_id"`
}
type CheckOldPointResponse struct {
	Meta Meta           `json:"meta"`
	Data *loyalty.Point `json:"data"`
}

type SearchPointRequestOld struct {
	Term              string  `form:"term"`
	UserId            string  `form:"user_id"`
	PointType         string  `form:"point_type"`
	Point             int64   `form:"point"`
	Source            string  `form:"source"`
	ForTransactionId  string  `form:"for_transaction_id"`
	TransactionAmount int64   `form:"transaction_amount"`
	Notes             string  `form:"notes"`
	CreatedFrom       string  `form:"created_from"`
	CreatedTo         string  `form:"created_to"`
	CreatedYmd        string  `form:"created_ymd"`
	Status            int32   `form:"status"`
	AppId             string  `form:"app_id"`
	PromotionPercent  float64 `form:"promotion_percent"`
	CampaignCode      string  `form:"campaign_code"`
	Channel           string  `form:"channel"`
	Rate              float64 `form:"rate"`
	JsonDetail        string  `form:"json_detail"`
	PageIndex         int32   `form:"page_index"`
	PageSize          int32   `form:"page_size"`
}
type SearchPointRequest struct {
	Term              string `form:"term"`
	UserId            string `form:"user_id"`
	PointType         string `form:"point_type"`
	Point             string `form:"point"`
	Source            string `form:"source"`
	ForTransactionId  string `form:"for_transaction_id"`
	TransactionAmount string `form:"transaction_amount"`
	Notes             string `form:"notes"`
	CreatedFrom       string `form:"created_from"`
	CreatedTo         string `form:"created_to"`
	CreatedYmd        string `form:"created_ymd"`
	Status            string `form:"status"`
	AppId             string `form:"app_id"`
	PromotionPercent  string `form:"promotion_percent"`
	CampaignCode      string `form:"campaign_code"`
	Channel           string `form:"channel"`
	Rate              string `form:"rate"`
	JsonDetail        string `form:"json_detail"`
	SortColumn        string `form:"sort_column"`
	SortDirection     string `form:"sort_direction"`
	PageIndex         int32  `form:"page_index"`
	PageSize          int32  `form:"page_size"`
}

type SearchPointResponse struct {
	Meta Meta         `json:"meta"`
	Data []*PointInfo `json:"data"`
	TotalRecord int64 `json:"total_record"`
}
