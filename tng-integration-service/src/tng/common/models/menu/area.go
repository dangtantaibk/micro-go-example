package menu

type Area struct {
	ID int64 `orm:"column(area_id)"`
	AreaName string `orm:"column(area_name); size(32)"`
	Description string `orm:"column(description)"`
	Status int `orm:"column(status)"`
}
