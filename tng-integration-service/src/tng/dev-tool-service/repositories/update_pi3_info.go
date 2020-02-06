package repositories

import (
	"github.com/astaxie/beego/orm"
	"tng/common/models/dev-tool"
	"tng/common/utils/db"
)

// Pi3UpdateInfoRepository represents a repository.
type Pi3UpdateInfoRepository interface {
	SetPi3UpdateInfo(ormer *db.DB, info *dev_tool.DeviceUpdateInfo) error
	GetPi3UpdateInfo(ormer *db.DB, posID string) (*dev_tool.DeviceUpdateInfo, error)
	DelPi3UpdateInfo(ormer *db.DB, posID string) error
}

type pi3UpdateInfoRepository struct{}

// NewPi3UpdateInfoRepository create a new instance Repository.
func NewPi3UpdateInfoRepository(dbFactory db.Factory, ) Pi3UpdateInfoRepository {
	return &pi3UpdateInfoRepository{}
}

func (h *pi3UpdateInfoRepository) SetPi3UpdateInfo(ormer *db.DB, info *dev_tool.DeviceUpdateInfo) error {
	_, err := ormer.Insert(info)
	return err
}

func (h *pi3UpdateInfoRepository) GetPi3UpdateInfo(ormer *db.DB, posID string) (*dev_tool.DeviceUpdateInfo, error) {
	cond := orm.NewCondition().
		And("pos_id", posID)
	cfg := &dev_tool.DeviceUpdateInfo{}
	qs := ormer.QueryTable(new(dev_tool.DeviceUpdateInfo))
	err := qs.SetCond(cond).One(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (h *pi3UpdateInfoRepository) DelPi3UpdateInfo(ormer *db.DB, posID string) error {
	_, err := ormer.QueryTable(new(dev_tool.DeviceUpdateInfo)).
		Filter("pos_id", posID).Delete()
	if err != nil {
		return err
	}
	return nil
}
