package dtos

type PromotionRequest struct {
	AppID     string            `json:"app_id"`
	Data      *TixPromotionData `json:"data"`
	UserAgent string            `json:"user_agent"`
	ClientIP  string            `json:"client_ip"`
	TS        int64             `json:"ts"`
	CheckSum  string            `json:"check_sum"`
}

type PromotionResponse struct {
	Meta Meta                      `json:"meta"`
	Data *TixPromotionDataResponse `json:"data"`
}
