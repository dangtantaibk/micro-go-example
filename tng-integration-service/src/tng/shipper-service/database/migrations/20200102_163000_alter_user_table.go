package migrations


import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AlterUserTable_20200102_163000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterUserTable_20200102_163000{}
	m.Created = "20200102_163000"
	migration.Register("UserTable_20200102_163000", m)
}

// Run the migrations
func (m *AlterUserTable_20200102_163000) Up() {
	m.SQL(`alter table user add password varchar(64) null;`)
	m.SQL(`create unique index user_phone_number_uindex on user (phone_number);`)
	m.SQL(`alter table user modify social_id bigint default 0 null;`)
	m.SQL(`drop index social_id on user;`)

}

// Reverse the migrations
func (m *AlterUserTable_20200102_163000) Down() {
}