package cron_job

type Item struct {
	ID               int64   `orm:"column(item_id)"`
	ItemName         string  `orm:"column(item_name);size(64)"`
	ItemCode         string  `orm:"column(item_code);size(32)"`
	BarCode          string  `orm:"column(barcode);size(32)"`
	Price            float64 `orm:"column(price)"`
	Order            int     `orm:"column(order)"`
	UnitID           int     `orm:"column(unit_id)"`
	CategoryID       int     `orm:"column(category_id)"`
	CateMask         int     `orm:"column(cate_mask)"`
	PrinterMask      int     `orm:"column(printer_mask)"`
	KitchenAreaID    int     `orm:"column(kitchen_area_id)"`
	PromotionID      int     `orm:"column(promotion_id)"`
	ImgPath          string  `orm:"column(img_path);size(256)"`
	ImgCrc           string  `orm:"column(img_crc);size(64)"`
	Description      string  `orm:"column(description);size(256)"`
	Inventory        int     `orm:"column(inventory)"`
	ModifiedBy       string  `orm:"column(modified_by);size(64)"`
	ModifiedDateTime string  `orm:"column(modified_date_time)"`
	Status           int     `orm:"column(status)"`
	MarkedDelete     int     `orm:"column(marked_delete)"`
	OriginalPrice    float64 `orm:"column(original_price)"`
	PromotionPrice   float64 `orm:"column(promotion_price)"`
	AutoSync         int     `orm:"column(auto_sync)"`
}
