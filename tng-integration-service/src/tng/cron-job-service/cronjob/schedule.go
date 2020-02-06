package cronjob

/*
* c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour") })
* c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
* c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
* c.AddFunc("@hourly",      func() { fmt.Println("Every hour, starting an hour from now") })
* c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
*/

var (
	// ScheduleDayByDayMenuSchedule start at 1h:0m
	ScheduleDayByDayMenuSchedule = "0 1 * * *"
	ScheduleDayByDayUpdateMenuSchedule = "0 1 * * *"
)
