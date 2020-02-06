package shiper

type InvoiceDetail struct {
	ID            int64  `orm:"column(invoice_detail_id)"`
	InvoiceID     int64  `orm:"column(invoice_id)"`
	ItemID        int64  `orm:"column(item_id)"`
	ItemName      string `orm:"column(item_name)"`
	Quantity      int    `orm:"column(quantity)"`
	UnitPrice     int64  `orm:"column(unit_price)"`
	Amount        int64  `orm:"column(amount)"`
	OriginalPrice int64  `orm:"column(original_price)"`
	PromotionID   int64  `orm:"column(promotion_id)"`
	StaffID       int    `orm:"column(staff_id)"`
	Note          string `orm:"column(note)"`
}
