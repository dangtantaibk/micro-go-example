package dtos

type OrderInfo struct {
	AppUser     string `json:"app_user"`
	AppTime     int64  `json:"app_time"`
	Amount      int64  `json:"amount"`
	AppTransID  string `json:"app_trans_id"`
	EmbedData   string `json:"embed_data"`
	Item        string `json:"item"`
	Description string `json:"description"`
	Hmac        string `json:"hmac"`
}
