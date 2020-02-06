package shiper

type General struct {
	ID    int64  `orm:"column(id)"`
	Name  string `orm:"column(name)"`
	Value string `orm:"column(value)"`
}
