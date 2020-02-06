package dtos

type DataPi3UpdateInfo struct {
	URL string `json:"url"`
}

type GetPi3UpdateInfoRequest struct {
	PosID     string `form:"pos_id"`
	TimeStamp int64  `form:"timestamp"`
	Sig       string `form:"sig"`
}

type GetPi3UpdateInfoResponse struct {
	Meta Meta               `json:"meta"`
	Data *DataPi3UpdateInfo `json:"data"`
}

type RegisterPi3Request struct {
	PosID      string `json:"pos_id"`
	DeviceType int    `json:"device_type"`
	TimeStamp  int64  `json:"timestamp"`
	Sig        string `json:"sig"`
}

type RegisterPi3Response struct {
	Meta Meta `json:"meta"`
}
