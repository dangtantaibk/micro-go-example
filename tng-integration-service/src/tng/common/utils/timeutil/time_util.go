package timeutil

import (
	"time"
)

// ParseTime converts string to time object
func ParseTime(t, layout string) *time.Time {
	if t != "" {
		t, err := time.Parse(layout, t)
		if err != nil {
			return nil
		}
		return &t
	}
	return nil
}

// ConvertDateToString will format date to string according to specified layout
func ConvertDateToString(date *time.Time, layout string) string {

	if date == nil {
		return ""
	}

	return date.Format(layout)
}
