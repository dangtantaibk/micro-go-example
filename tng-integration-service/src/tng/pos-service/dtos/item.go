package dtos

type Item struct {
	ItemID           int64  `json:"item_id"`
	ItemName         string `json:"item_name"`
	ItemCode         string `json:"item_code"`
	Price            int64  `json:"price"`
	UnitName         string `json:"unit_name"`
	CategoryName     string `json:"category_name"`
	CategoryID       int    `json:"category_id"`
	ImgPath          string `json:"img_path"`
	ImgCrc           string `json:"img_crc"`
	Description      string `json:"description"`
	Inventory        int    `json:"inventory"`
	Status           int    `json:"status"`
	CreateBy         string `json:"create_by"`
	CreateDate       string `json:"create_date"`
	BarCode          string `json:"bar_code"`
	ModifiedBy       string `json:"modified_by"`
	ModifiedDateTime string `json:"modified_date_time"`
	Order            int    `json:"order"`
	OriginalPrice    int64  `json:"original_price"`
	PromotionType    int    `json:"promotion_type"`
	PromotionID      int64  `json:"promotion_id"`
	CateMask         int64  `json:"cate_mask"`
	PrinterMask      int    `json:"printer_mask"`
	KitchenAreaID    int64  `json:"kitchen_area_id"`
	AreaID           int    `json:"area_id"`
	DayOfSchedule    int    `json:"day_of_schedule"`
	BeginTime        string `json:"begin_time"`
	EndTime          string `json:"end_time"`
	InStock          int    `json:"in_stock"`
	PromotionPrice   int64  `json:"promotion_price"`
	MerchantCode     string `json:"merchant_code"`
}
