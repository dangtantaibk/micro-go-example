package zalopay

const (
	EndpointCreateOrder     = "/v001/tpe/createorder"
	EndpointGetOrderStatus  = "/v001/tpe/getstatusbyapptransid"
	EndpointRefund          = "/v001/tpe/partialrefund"
	EndpointGetRefundStatus = "/v001/tpe/getpartialrefundstatus"
	ZlpBankCode             = "zalopayapp"
)

type ZlpCreateOrderRequest struct {
	MerchantCode string `json:"merchantcode"`
	AppID        string `json:"appid"`
	AppTransID   string `json:"apptransid"`
	AppUser      string `json:"appuser"`
	AppTime      string `json:"apptime"`
	Amount       string `json:"amount"`
	EmbedData    string `json:"embeddata"`
	Item         string `json:"item"`
	BankCode     string `json:"bankcode"`
	Mac          string `json:"mac"`
	Description  string `json:"description"`
}

type ZlpCreateOrderResponse struct {
	ReturnCode    int    `json:"returncode"`
	ReturnMessage string `json:"returnmessage"`
	OrderUrl      string `json:"orderurl"`
	ZpTransToken  string `json:"zptranstoken"`
}

type ZlpGetInvoiceStatusRequest struct {
	MerchantCode string `json:"merchantcode"`
	AppID        string `json:"appid"`
	AppTransID   string `json:"apptransid"`
}

type ZlpGetInvoiceStatusResponse struct {
	ReturnCode    int    `json:"returncode"`
	ReturnMessage string `json:"returnmessage"`
	IsProcessing  bool   `json:"isprocessing"`
	Amount        int64  `json:"amount"`
	ZpTransID     int64  `json:"zptransid"`
}

type ZlpRefundInvoiceRequest struct {
	MerchantCode string `json:"merchantcode"`
	MRefundID    string `json:"mrefundid"`
	AppID        string `json:"appid"`
	ZpTransID    string `json:"zptransid"`
	Amount       string `json:"amount"`
	Description  string `json:"description"`
}

type ZlpRefundInvoiceResponse struct {
	ReturnCode    int    `json:"returncode"`
	ReturnMessage string `json:"returnmessage"`
	RefundID      string `json:"refundid"`
}

type ZlpGetRefundStatusRequest struct {
	MerchantCode string `json:"merchantcode"`
	AppID        string `json:"appid"`
	MRefundID    string `json:"mrefundid"`
	TimeStamp    int64  `json:"timestamp"`
}

type ZlpGetRefundStatusResponse struct {
	ReturnCode    int    `json:"returncode"`
	ReturnMessage string `json:"returnmessage"`
}
