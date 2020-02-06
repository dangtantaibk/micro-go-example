package dtos

type WarmUpScheduleRequest struct {
	MerchantCode string `json:"merchant_code"` // if merchant code isEmpty -> get all
	TimeStamp    int64  `json:"time_stamp"`
	Sig          string `json:"sig"`
}

type WarmUpScheduleResponse struct {
	Meta Meta `json:"meta"`
}

type Schedule struct {
	Title         string  `json:"title"`
	AreaID        int     `json:"area_id"`
	ItemID        int     `json:"item_id"`
	BeginTime     string  `json:"begin_time"`
	EndTime       string  `json:"end_time"`
	Price         float64 `json:"price"`
	OrderIng      int     `json:"order_ing"`
	InStock       int     `json:"in_stock"`
	Published     int     `json:"published"`
	Created       string  `json:"created"`
	CreatedBy     string  `json:"created_by"`
	Modified      string  `json:"modified"`
	ModifiedBy    string  `json:"modified_by"`
	JsonString    string  `json:"json_string"`
	DayOfSchedule int     `json:"day_of_schedule"`
}