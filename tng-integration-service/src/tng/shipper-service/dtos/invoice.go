package dtos

type Invoice struct {
	ID                  int64   `json:"invoice_id"`
	InvoiceCode         string  `json:"invoice_code"`
	InvoiceNo           int     `json:"invoice_no"`
	Amount              float64 `json:"amount"`
	ExpectedAmount      float64 `json:"expected_amount"`
	Discount            float64 `json:"discount"`
	PromotionID         int     `json:"promotion_id"`
	CreatedDateTime     string  `json:"created_date_time"`
	PaymentStatus       int     `json:"payment_status"`
	PaymentMethod       int     `json:"payment_method"`
	PaymentDateTime     string  `json:"payment_date_time"`
	ZpTransID           string  `json:"zp_trans_id"`
	ZpServerTime        int64   `json:"zp_server_time"`
	ZpTransToken        string  `json:"zp_trans_token"`
	ZpUserID            string  `json:"zp_user_id"`
	AuditStatus         int     `json:"audit_status"`
	AuditBy             string  `json:"audit_by"`
	AuditDateTime       string  `json:"audit_date_time"`
	StaffID             int     `json:"staff_id"`
	MachineName         string  `json:"machine_name"`
	AreaID              int     `json:"area_id"`
	UserID              int64   `json:"user_id"`
	Note                string  `json:"note"`
	Printed             int     `json:"printed"`
	VposToken           string  `json:"vpos_token"`
	DeliveryPhoneNumber string  `json:"delivery_phone_number"`
	DeliveryAddress     string  `json:"delivery_address"`
	DeliveryNote        string  `json:"delivery_note"`
	OrderMethod         int     `json:"order_method"`
	ShipperName         string  `json:"shipper_name"`
	ShipperPhone        string  `json:"shipper_phone"`
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

type ListInvoiceRequest struct {
	Date              string `form:"date"`
	MerchantCode      string `form:"merchant_code"`
	CurrentPage       int32  `form:"current_page"`
	TotalTransPerPage int32  `form:"total_trans_per_page"`
}

type ListInvoiceResponse struct {
	Meta           Meta      `json:"meta"`
	TotalRecord    int64     `json:"total_record"`
	TotalAllRecord int64     `json:"total_all_record"`
	Data           []Invoice `json:"data"`
}

type UpdateInvoiceStatusRequest struct {
	InvoiceCode   string `json:"invoice_code"`
	PaymentStatus int    `json:"payment_status"`
	MerchantCode  string `json:"merchant_code"`
}

type UpdateInvoiceStatusResponse struct {
	Meta Meta `json:"meta"`
}

type ScanQRCodeRequest struct {
	MerchantCode string `form:"merchant_code"`
	VposToken    string `form:"vpos_token"`
}

type ScanQRCodeResponse struct {
	Meta          Meta            `json:"meta"`
	Invoice       Invoice         `json:"invoice"`
	InvoiceDetail []*InvoiceDetail `json:"invoice_detail"`
}

type GetInvoiceDetailRequest struct {
	MerchantCode string `form:"merchant_code"`
	InvoiceCode  string `form:"invoice_code"`
}

type GetInvoiceDetailResponse struct {
	Meta          Meta            `json:"meta"`
	Invoice       Invoice         `json:"invoice"`
	InvoiceDetail []*InvoiceDetail `json:"invoice_detail"`
}
