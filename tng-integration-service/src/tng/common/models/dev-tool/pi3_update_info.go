package dev_tool

type DeviceUpdateInfo struct {
	ID             int64  `orm:"column(id)"`
	MerchantID     int    `orm:"column(merchant_id)"`
	AppID          int    `orm:"column(app_id)"`
	StoreID        int    `orm:"column(store_id)"`
	DeviceType     int    `orm:"column(device_type)"`
	PosID          string `orm:"column(pos_id);size(255)"`
	UrlFileExecute string `orm:"column(url_file_execute);size(255)"`
}
