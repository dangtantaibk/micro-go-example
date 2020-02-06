package pos

type InvoiceDetail struct {
	ID            int64   `orm:"column(invoice_detail_id)"`
	InvoiceID     int64   `orm:"column(invoice_id)"`
	ItemID        int     `orm:"column(item_id)"`
	ItemName      string  `orm:"column(item_name);size(64)"`
	Quantity      int     `orm:"column(quantity)"`
	UnitPrice     float64 `orm:"column(unit_price)"`
	Amount        float64 `orm:"column(amount)"`
	OriginalPrice float64 `orm:"column(original_price)"`
	PromotionID   int     `orm:"column(promotion_id)"`
	StaffID       int     `orm:"column(staff_id)"`
	Note          string  `orm:"column(note);size(256)"`
}
