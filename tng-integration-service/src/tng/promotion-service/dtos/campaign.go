package dtos

import "time"

type Campaign struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	From            string `json:"from"`
	To              string `json:"to"`
	RepeatType      int    `json:"repeat_type"`
	RepeatValue     string `json:"repeat_value"`
	ExcludeDays     string `json:"exclude_days"`
	Status          int    `json:"status"`
	TimeFrom        string `json:"time_from"`
	TimeTo          string `json:"time_to"`
	AppID           int    `json:"app_id"`
	MinOrderAmt     string `json:"min_order_amt"`
	MaxOrderAmt     string `json:"max_order_amt"`
	MaxAmt          int    `json:"max_amt"`
	MaxTransaction  int    `json:"max_transaction"`
	PromoType       int    `json:"promo_type"`
	SchemeType      int    `json:"scheme_type"`
	DiscountPercent string `json:"discount_percent"`
	DiscountAmt     string `json:"discount_amt"`
	FixedDiscount   int    `json:"fixed_discount"`
	Created         string `json:"created"`
	CreatedBy       string `json:"created_by"`
	Modified        string `json:"modified"`
	ModifiedBy      string `json:"modified_by"`
	Description     string `json:"description"`
	Json            string `json:"json"`
	UCode           string `json:"u_code"`
	ENV             string `json:"env"`
	Channel         string `json:"channel"`
	PaymentMethod   string `json:"payment_method"`
	PaymentChannel  string `json:"payment_channel"`
	BankCode        string `json:"bank_code"`
}

type InsertCampaignRequest struct {
	Title           string    `json:"title"`
	From            time.Time `json:"from"`
	To              time.Time `json:"to"`
	RepeatType      int       `json:"repeat_type"`
	RepeateValue    string    `json:"repeate_value"`
	ExcludeDays     string    `json:"exclude_days"`
	Status          int       `json:"status"`
	TimeFrom        string    `json:"time_from"`
	TimeTo          string    `json:"time_to"`
	AppID           int       `json:"app_id"`
	MinOrderAmt     float64   `json:"min_order_amt"`
	MaxOrderAmt     float64   `json:"max_order_amt"`
	MaxAmt          int64     `json:"max_amt"`
	MaxTransaction  int       `json:"max_transaction"`
	PromoType       int       `json:"promo_type"`
	SchemeType      int       `json:"scheme_type"`
	DiscountPercent float64   `json:"discount_percent"`
	DiscountAmt     float64   `json:"discount_amt"`
	FixedDiscount   int       `json:"fixed_discount"`
	Description     string    `json:"description"`
	Json            string    `json:"json"`
	UCode           string    `json:"ucode"`
	TimeStamp       int64     `json:"time_stamp"`
	Sig             string    `json:"sig"`
}

type InsertCampaignResponse struct {
	Meta Meta `json:"meta"`
}

type ListCampaignRequest struct {
	AppID     string `form:"appid"`
	PageIndex int32  `form:"page_index"`
	PageSize  int32  `form:"page_size"`
}

type ListCampaignResponse struct {
	Meta Meta        `json:"meta"`
	Data []*Campaign `json:"data"`
}
