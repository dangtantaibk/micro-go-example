package h5_intergration

type H5ZaloPay struct {
	ID         int64  `orm:"column(id)"`
	AppID      int    `orm:"column(app_id)"`
	TokenType  int    `orm:"column(token_type)"`
	MAToken    string `orm:"column(ma_token);size(255)"`
	MBToken    string `orm:"column(mb_token);size(255)"`
	Status     int    `orm:"column(status)"`
	ExpireTime int64  `orm:"column(expire_time)"`
	UserID     string `orm:"column(user_id);size(255)"`
}