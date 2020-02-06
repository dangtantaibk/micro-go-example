package location

import (
	"time"
)

var (
	secondsOfVN = int((7 * time.Hour).Seconds())
	VNLocation  = time.FixedZone("Vietnam", secondsOfVN)
	HCMLocation = "Asia/Ho_Chi_Minh"
)

func GetVNCurrentTime() *time.Time {
	t := time.Now().In(VNLocation)
	return &t
}

func GetVNCurrentTimeYMD(year, month, day int) *time.Time {
	t := time.Now().In(VNLocation).AddDate(year, month, day)
	return &t
}
