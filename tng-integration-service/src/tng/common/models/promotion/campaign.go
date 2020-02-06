package promotion

import "time"

type Campaign struct {
	ID              int64     `orm:"column(id)"`
	Title           string    `orm:"column(title);size(256)"`
	From            time.Time `orm:"column(from);type(datetime)"`
	To              time.Time `orm:"column(to);type(datetime)"`
	RepeatType      int       `orm:"column(repeat_type)"`
	RepeateValue    string    `orm:"column(repeate_value)"`
	ExcludeDays     string    `orm:"column(exclude_days)"`
	Status          int       `orm:"column(status)"`
	TimeFrom        string    `orm:"column(time_from);size(8)"`
	TimeTo          string    `orm:"column(time_to);size(8)"`
	AppID           int       `orm:"column(appid)"`
	MinOrderAmt     float64   `orm:"column(min_order_amt)"`
	MaxOrderAmt     float64   `orm:"column(max_order_amt)"`
	MaxAmt          int64     `orm:"column(max_amt)"`
	MaxTransaction  int       `orm:"column(max_transaction)"`
	PromoType       int       `orm:"column(promo_type)"`
	SchemeType      int       `orm:"column(scheme_type)"`
	DiscountPercent float64   `orm:"column(discount_percent)"`
	DiscountAmt     float64   `orm:"column(discount_amt)"`
	FixedDiscount   int       `orm:"column(fixed_discount)"`
	Created         time.Time `orm:"column(created);type(datetime)"`
	CreatedBy       string    `orm:"column(created_by);size(45)"`
	Modified        time.Time `orm:"column(modified);type(datetime)"`
	ModifiedBy      string    `orm:"column(modified_by);size(45)"`
	Description     string    `orm:"column(description)"`
	Json            string    `orm:"column(json)"`
	UCode           string    `orm:"column(ucode);size(45)"`
}
