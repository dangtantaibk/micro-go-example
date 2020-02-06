package h5_zalopay

const (
	EndpointGetMBToken = "/colossus/api/v1/grantedmbtoken"
	EndpointZPOrderURL = "/colossus/api/v1/paymenturl/orderurl"
)

type DataH5ZaloPayOrderURL struct {
	PaymentURL string `json:"url"`
}

type DataH5ZaloPayGrantedMBToken struct {
	MBToken string `json:"mbtoken"`
}

type GetGrantedMBTokenResp struct {
	Data          *DataH5ZaloPayGrantedMBToken `json:"data"`
	ReturnCode    int                          `json:"returncode"`
	ReturnMessage string                       `json:"returnmessage"`
	TraceID       string                       `json:"traceid"`
}

type H5ZaloPayOrderURLReq struct {
	ZPOrderURL string `json:"zporderurl"`
	AppID      int    `json:"appid"`
	MBToken    string `json:"mbtoken"`
	ClientID   int64  `json:"clientid"`
	ReqDate    int64  `json:"reqdate"`
	Sig        string `json:"sig"`
}

type H5ZaloPayOrderURLResp struct {
	Data          *DataH5ZaloPayOrderURL `json:"data"`
	ReturnCode    int                    `json:"returncode"`
	ReturnMessage string                 `json:"returnmessage"`
}
