package shiper

type Shipper struct {
	ID       int64  `orm:"column(id)"`
	Name     string `orm:"column(name);size(128)"`
	Phone    string `orm:"column(phone);size(128)"`
	Password string `orm:"column(password);size(64)"`
	Status   int   	`orm:"column(status);size(1)"`
}
