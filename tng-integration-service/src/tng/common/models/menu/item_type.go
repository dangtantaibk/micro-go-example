package menu


type ItemType struct {
	ID int64 `orm:"column(item_type_id)"`
	ItemTypeName string `orm:"column(item_type_name); size(64)"`
	Status int `orm:"column(status)"`
}
