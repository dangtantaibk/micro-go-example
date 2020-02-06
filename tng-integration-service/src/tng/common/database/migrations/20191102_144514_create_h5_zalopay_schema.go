package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateH5ZaloPaySchema_20191102_144514 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateH5ZaloPaySchema_20191102_144514{}
	m.Created = "20191102_144514"

	migration.Register("CreateH5ZaloPaySchema_20191102_144514", m)
}

// Run the migrations
func (m *CreateH5ZaloPaySchema_20191102_144514) Up() {
	m.SQL(`
        CREATE TABLE IF NOT EXISTS h5_zalo_pay (
            id 			bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
			app_id 		int(11) NOT NULL DEFAULT 0,
			token_type 	int(2) NOT NULL DEFAULT -1,
			ma_token 	varchar(255) NOT NULL DEFAULT '',
			mb_token 	varchar(255) NOT NULL DEFAULT '',
			status 		int(2) NOT NULL DEFAULT -1,
			expire_time bigint NOT NULL DEFAULT 0,
            user_id 	varchar(255) NOT NULL DEFAULT ''
        ) ENGINE=InnoDB;
    `)
}

// Reverse the migrations
func (m *CreateH5ZaloPaySchema_20191102_144514) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS h5_zalo_pay")
}
