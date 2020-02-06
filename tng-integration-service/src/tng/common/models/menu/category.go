package menu

import "time"

type Category struct {
	ID int64 `orm:"column(category_id)"`
	CategoryName string `orm:"column(category_name); size(32)"`
	CategoryValue int64 `orm:"column(category_value)"`
	CategoryType int64 `orm:"column(category_type)"`
	Order int `orm:"column(order)"`
	Status int `orm:"column(status)"`
	AreaID int64 `orm:"column(area_id)"`
	ModifiedBy string `orm:"column(modified_by)"`
	ModifiedDateTime time.Time `orm:"column(modified_date_time)"`
}