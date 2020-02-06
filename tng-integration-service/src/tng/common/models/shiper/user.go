package shiper

type User struct {
	ID          int64  `orm:"column(user_id)"`
	SocialName  string `orm:"column(social_name);size(32)"`
	SocialId    int64  `orm:"column(social_id)"`
	Address     string `orm:"column(address);size(500)"`
	PhoneNumber string `orm:"column(phone_number);size(15)"`
	Password    string `orm:"column(password);size(64)"`
}
