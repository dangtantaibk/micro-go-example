package pos

type Invoice struct {
	ID              int64   `orm:"column(invoice_id)"`
	InvoiceCode     string  `orm:"column(invoice_code);size(32)"`
	InvoiceNo       int     `orm:"column(invoice_no)"`
	Amount          float64 `orm:"column(amount)"`
	ExpectedAmount  float64 `orm:"column(expected_amount)"`
	Discount        float64 `orm:"column(discount)"`
	PromotionID     int     `orm:"column(promotion_id)"`
	CreatedDateTime string  `orm:"column(created_date_time)"`
	PaymentStatus   int     `orm:"column(payment_status)"`
	PaymentMethod   int     `orm:"column(payment_method)"`
	PaymentDateTime string  `orm:"column(payment_date_time)"`
	ZpTransID       string  `orm:"column(zptransid);size(32)"`
	ZpServerTime    int64   `orm:"column(zpservertime)"`
	ZpTransToken    string  `orm:"column(zptranstoken);size(64)"`
	ZpUserID        string  `orm:"column(zpuserid);size(64)"`
	AuditStatus     int     `orm:"column(audit_status)"`
	AuditBy         string  `orm:"column(audit_by);size(64)"`
	AuditDateTime   string  `orm:"column(audit_date_time)"`
	StaffID         int     `orm:"column(staff_id)"`
	MachineName     string  `orm:"column(machine_name);size(16)"`
	AreaID          int     `orm:"column(area_id)"`
	UserID          int64   `orm:"column(user_id)"`
	Note            string  `orm:"column(note);size(128)"`
	Printed         int     `orm:"column(printed)"`
	VposToken       string  `orm:"column(vpostoken);size(64)"`
}
