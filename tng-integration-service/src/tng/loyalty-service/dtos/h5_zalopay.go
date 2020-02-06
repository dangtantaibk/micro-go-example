package dtos

const (
	TOKEN_STT_UNKNOWN = 0
	TOKEN_STT_ENABLE  = 1
	TOKEN_STT_DISABLE = 2
)
const (
	TOKEN_TYPE_UNKNOWN    = 0
	TOKEN_TYPE_H5_ZALOPAY = 1
)

type DataH5ZaloPayOrderURL struct {
	PaymentURL string `json:"paymenturl"`
}

type GetH5ZaloPayRequest struct {
	AppID     int    `form:"app_id"`
	UserID    string `form:"user_id"`
	MAToken   string `form:"ma_token"`
	TimeStamp int64  `form:"timestamp"`
	Sig       string `form:"sig"`
}

type GetH5ZaloPayResp struct {
	Meta Meta           `json:"meta"`
}

type H5ZaloPayOrderURLReq struct {
	UserID     string `json:"user_id"`
	ZPOrderURL string `json:"zporderurl"`
	AppID      int    `json:"app_id"`
	MAToken    string `json:"ma_token"`
	TimeStamp  int64  `json:"timestamp"`
	Sig        string `json:"sig"`
}

type H5ZaloPayOrderURLResp struct {
	Meta Meta                   `json:"meta"`
	Data *DataH5ZaloPayOrderURL `json:"data"`
}

type TokenInfo struct {
	AppID      int    `json:"app_id,omitempty"`
	TokenType  int    `json:"token_type,omitempty"`
	MAToken    string `json:"ma_token"`
	MBToken    string `json:"mb_token"`
	ExpireTime int64  `json:"expire_time,omitempty"`
	Status     int    `json:"status,omitempty"`
	UserID     string `json:"user_id"`
}
