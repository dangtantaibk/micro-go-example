package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateDeviceDeviceUpdateSchema_20191102_144514 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateDeviceDeviceUpdateSchema_20191102_144514{}
	m.Created = "20191102_144514"

	migration.Register("CreateDeviceDeviceUpdateSchema_20191102_144514", m)
}

// Run the migrations
func (m *CreateDeviceDeviceUpdateSchema_20191102_144514) Up() {
	m.SQL(`
        CREATE TABLE IF NOT EXISTS device_update_info (
            id 					bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
			merchant_id 		int(11) NOT NULL DEFAULT 0,
			app_id 				int(11) NOT NULL DEFAULT 0,
			store_id 			int(11) NOT NULL DEFAULT 0,
			pos_id	 			varchar(255) NOT NULL DEFAULT '',
			device_type 		int(2) NOT NULL DEFAULT -1,
			url_file_execute 	varchar(255) NOT NULL DEFAULT '',
			create_time 		bigint NOT NULL DEFAULT 0,
			update_time 		bigint NOT NULL DEFAULT 0
        ) ENGINE=InnoDB;
    `)
}

// Reverse the migrations
func (m *CreateDeviceDeviceUpdateSchema_20191102_144514) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS device_update_info")
}
