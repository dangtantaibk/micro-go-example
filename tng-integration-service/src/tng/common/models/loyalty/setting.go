package loyalty

type Setting struct {
	ID                   int32  `orm:"column(input_money_per_point)"`
	OutputMoneyPerPoint  int32  `orm:"column(output_money_per_point)"`
	PeriodOfClassByMonth int32  `orm:"column(period_of_class_by_month)"`
	JsonDetail           string `orm:"column(json_detail)"`
}
