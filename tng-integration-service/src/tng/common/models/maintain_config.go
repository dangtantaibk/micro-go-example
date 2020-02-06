package models

// MaintainStatus const definition.
const (
	MaintainStatusEnable  = "ENABLE"
	MaintainStatusDisable = "DISABLE"
)

// MaintainConfiguration is table contains config for maintain modules.
type MaintainConfiguration struct {
	BaseModel
	Module string `orm:"column(module);size(100)"`
	Action string `orm:"column(action);size(100)"`
	Status string `orm:"column(status);size(10);default(ENABLE)"`
}
