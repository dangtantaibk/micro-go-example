package shiper

type Item struct {
	ID               int64   `orm:"column(item_id)"`
	ItemName         string  `orm:"column(item_name)"`
	ItemCode         string  `orm:"column(item_code)"`
	Barcode          string  `orm:"column(barcode)"`
	Price            float64 `orm:"column(price)"`
	Order            int16   `orm:"column(order)"`
	UnitId           int32   `orm:"column(unit_id)"`
	CategoryId       int64   `orm:"column(category_id)"`
	CateMask         int64   `orm:"column(cate_mask)"`
	PrinterMask      int64   `orm:"column(printer_mask)"`
	KitchenAreaId    int64   `orm:"column(kitchen_area_id)"`
	PromotionId      int64   `orm:"column(promotion_id)"`
	ImgPath          string  `orm:"column(img_path)"`
	ImgCrc           string  `orm:"column(img_crc)"`
	Description      string  `orm:"column(description)"`
	Inventory        int64   `orm:"column(inventory)"`
	ModifiedBy       string  `orm:"column(modified_by)"`
	ModifiedDateTime string  `orm:"column(modified_date_time)"`
	Status           int64   `orm:"column(status)"`
	MarkedDelete     int     `orm:"column(marked_delete)"`
	OriginalPrice    float64 `orm:"column(original_price)"`
	PromotionPrice   float64 `orm:"column(promotion_price)"`
	IsSideDish       int     `orm:"column(is_side_dish)"`
	ParentId         int64   `orm:"column(parent_id)"`
	ItemTypeId       int64   `orm:"column(item_type_id)"`
	AutoSync         int32   `orm:"column(auto_sync)"`
}
