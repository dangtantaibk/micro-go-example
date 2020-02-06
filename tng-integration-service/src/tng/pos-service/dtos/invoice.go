package dtos

type Invoice struct {
	InvoiceID       int64  `json:"invoice_id"`
	InvoiceCode     string `json:"invoice_code"`
	InvoiceNO       int64  `json:"invoice_no"`
	Amount          int64  `json:"amount"`
	AmountGross     int64  `json:"amount_gross"`
	CreatedDateTime string `json:"created_date_time"`
	PaymentStatus   int    `json:"payment_status"`
	PaymentMethod   int    `json:"payment_method"`
	ZPTransToken    string `json:"zptranstoken"`
	PaymentDateTime string `json:"payment_date_time"`
	StaffID         int64  `json:"staff_id"`
	StaffName       string `json:"staff_name"`
	Note            string `json:"note"`
	ZPTransID       string `json:"zptransid"`
	AuditStatus     int    `json:"audit_status"`
	AuditBy         string `json:"audit_by"`
	AuditDateTime   string `json:"audit_date_time"`
	MachineName     string `json:"machine_name"`
	MerchantName    string `json:"merchant_name"`
	MerchantCode    string `json:"merchant_code"`
	AreaID          int64  `json:"area_id"`
	AreaName        string `json:"area_name"`
	UserID          int64  `json:"user_id"`
	Printed         int    `json:"printed"`
	Used            int    `json:"used"`
	VposToken       string `json:"vpostoken"`
}

type InvoiceDetail struct {
	InvoiceDetailID int64  `json:"invoice_detail_id"`
	InvoiceID       int64  `json:"invoice_id"`
	ItemID          int64  `json:"item_id"`
	ItemName        string `json:"item_name"`
	ItemCode        string `json:"item_code"`
	Quantity        int    `json:"quantity"`
	Amount          int64  `json:"amount"`
	Price           int64  `json:"price"`
	OriginalPrice   int64  `json:"original_price"`
	PromotionType   int    `json:"promotion_type"`
	PromotionID     int64  `json:"promotion_id"`
	ImgPath         string `json:"img_path"`
	AreaID          int    `json:"area_id"`
}

type ItemInvoice struct {
	ItemID        int    `json:"item_id"`
	ItemName      string `json:"item_name"`
	Quantity      int    `json:"quantity"`
	Price         int64  `json:"price"`
	Amount        int64  `json:"amount"`
	OriginalPrice int64  `json:"original_price"`
	PromotionType int    `json:"promotion_type"`
}
type CreateInvoiceRequest struct {
	MerchantCode  string        `json:"merchant_code"`
	FoodsOrderID  string        `json:"foodsorder_id"`
	MachineName   string        `json:"machine_name"`
	Description   string        `json:"description"`
	AppUser       string        `json:"appuser"`
	Amount        string        `json:"amount"`
	DevID         string        `json:"devid"`
	Items         []*ItemInvoice `json:"items"`
	PaymentMethod int        `json:"payment_method"`
	PaymentCode   int           `json:"paymentcode"`
	AreaID        int           `json:"area_id"`
	Sig           string        `json:"sig"`
	IsWeb         bool          `json:"is_web"`
}

type DataInvoiceResponse struct {

}

type CreateInvoiceResponse struct {
	Meta Meta                `json:"meta"`
	Data DataInvoiceResponse `json:"data"`
}

type CancelInvoiceRequest struct {
	MerchantCode string `json:"merchant_code"`
	InvoiceCode  string `json:"invoice_code"`
}

type CancelInvoiceResponse struct {
	Meta Meta `json:"meta"`
}

type RefundInvoiceRequest struct {

}

type RefundInvoiceResponse struct {
	Meta Meta `json:"meta"`
}
