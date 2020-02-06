package dtos

type SettingInfo struct {
	ID                   int32  `json:"id"`
	OutputMoneyPerPoint  int32  `json:"output_money_per_point"`
	PeriodOfClassByMonth int32  `json:"period_of_class_by_month"`
	JsonDetail           string `json:"json_detail"`
}

type InsertOrUpdateSettingRequest struct {
	MetaRequest
	ID                   int32  `json:"id"`
	OutputMoneyPerPoint  int32  `json:"output_money_per_point"`
	PeriodOfClassByMonth int32  `json:"period_of_class_by_month"`
	JsonDetail           string `json:"json_detail"`
}

type InsertOrUpdateSettingResponse struct {
	Meta Meta `json:"meta"`
}

type DeleteSettingRequest struct {
	ID int32 `json:"id"`
}

type DeleteSettingResponse struct {
	Meta Meta `json:"meta"`
}

type ListSettingRequest struct {
	MetaRequest
	PageIndex int32        `form:"page_index"`
	PageSize  int32        `form:"page_size"`
}

type ListSettingResponse struct {
	Meta Meta           `json:"meta"`
	Data []*SettingInfo `json:"data"`
}

type GetSettingRequest struct {
	MetaRequest
	ID   int32 `form:"id"`
}

type GetSettingResponse struct {
	Meta Meta         `json:"meta"`
	Data *SettingInfo `json:"data"`
}
