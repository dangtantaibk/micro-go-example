package user_profile

type LoginType struct {
	ID    int64  `orm:"column(id);size(10)"`
	Title string `orm:"column(title);size(256)"`
}
