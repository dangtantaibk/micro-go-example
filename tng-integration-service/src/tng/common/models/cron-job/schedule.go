package cron_job

type Schedule struct {
	ID            int64   `orm:"column(id)"`
	Title         string  `orm:"column(title);size(256)"`
	AreaID        int     `orm:"column(area_id)"`
	ItemID        int     `orm:"column(item_id)"`
	BeginTime     string  `orm:"column(begin_time)"`
	EndTime       string  `orm:"column(end_time)"`
	Price         float64 `orm:"column(price)"`
	OrderIng      int     `orm:"column(ordering)"`
	InStock       int     `orm:"column(instock)"`
	Published     int     `orm:"column(published)"`
	Created       string  `orm:"column(created)"`
	CreatedBy     string  `orm:"column(created_by);size(45)"`
	Modified      string  `orm:"column(modified)"`
	ModifiedBy    string  `orm:"column(modified_by);size(45)"`
	JsonString    string  `orm:"column(json_string);size(256)"`
	DayOfSchedule int     `orm:"column(day_of_schedule)"`
}
