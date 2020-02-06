package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableMaintainConfiguration_20190507_151426 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableMaintainConfiguration_20190507_151426{}
	m.Created = "20190507_151426"

	migration.Register("CreateTableMaintainConfiguration_20190507_151426", m)
}

// Run the migrations
func (m *CreateTableMaintainConfiguration_20190507_151426) Up() {
	m.SQL(`
		CREATE TABLE IF NOT EXISTS maintain_configuration
		(
			id         bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
			created_at datetime              NOT NULL,
			updated_at datetime              NOT NULL,
			deleted_at datetime,
			module     varchar(100)          NOT NULL DEFAULT '',
			action     varchar(100)          NOT NULL DEFAULT '',
			status     varchar(10)           NOT NULL DEFAULT 'ENABLE'
		) ENGINE = InnoDB;
	`)
	m.SQL(`CREATE UNIQUE INDEX maintain_configuration_module_action_uindex ON maintain_configuration (module, action);`)
}

// Reverse the migrations
func (m *CreateTableMaintainConfiguration_20190507_151426) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
}
