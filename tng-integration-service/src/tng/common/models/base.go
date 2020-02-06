package models

import "time"

// Common date time format.
const (
	FormatYYYYMMDD        = "2006-01-02"
	FormatDateTime        = time.RFC3339
	FormatHHMMSSDDMMYY    = "15:04:05 02/01/2006"
	FormatYYMMDDHHMMSS    = "2006/01/02 15:04:05"
	FormatYYMMDD          = "060102"
	FormatYYYYMMDD1       = "20060102"
	FormatYYYMMDDHHMMSS   = "2006-01-02 15:04:05"
	FormatDefaultDateTime = "1970-01-01 00:00:00"
	DBDriver              = "mysql"
	DBOEAlias             = "oe"
	DBDefaultAlias        = "default"
	PreDatabase           = "vpos_"
)

// BaseModel is basic information all models in system.
type BaseModel struct {
	ID        int64      `orm:"column(id)"`
	CreatedAt time.Time  `orm:"column(created_at);type(datetime);auto_now_add"`
	UpdatedAt time.Time  `orm:"column(updated_at);type(datetime);auto_now"`
	DeletedAt *time.Time `orm:"column(deleted_at);type(datetime);null"`
}
