package repositories

import (
	"context"
	"github.com/astaxie/beego/orm"
	"tng/common/models/loyalty"
	"tng/common/utils/db"
)

type SettingRepository interface {
	InsertOrUpdate(context.Context, *db.DB, *loyalty.Setting) error
	Delete(context.Context, *db.DB, int32) error
	GetByID(context.Context, *db.DB, int32) (*loyalty.Setting, error)
	List(context.Context, *db.DB, int32, int32) ([]*loyalty.Setting, error)
}

type settingRepository struct{}

func NewSettingRepository(dbFactory db.Factory) SettingRepository {
	return &settingRepository{}
}

func (r *settingRepository) InsertOrUpdate(ctx context.Context, ormer *db.DB, dataInput *loyalty.Setting) error {
	_, err := ormer.InsertOrUpdate(dataInput)
	return err
}

func (r *settingRepository) Delete(ctx context.Context, ormer *db.DB, id int32) error {
	qs := ormer.QueryTable(new(loyalty.Setting))
	cond := orm.NewCondition().
		And("input-money_per_point", id)
	_, err := qs.SetCond(cond).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r *settingRepository) GetByID(ctx context.Context, ormer *db.DB, id int32) (*loyalty.Setting, error) {
	cond := orm.NewCondition().
		And("input_money_per_point", id)
	setting := &loyalty.Setting{}
	qs := ormer.QueryTable(new(loyalty.Setting))
	err := qs.SetCond(cond).One(setting)
	if err != nil {
		return nil, err
	}
	return setting, nil

}

func (r *settingRepository) List(ctx context.Context, ormer *db.DB, pageIndex int32, pageSize int32) ([]*loyalty.Setting, error) {
	var (
		list       []*loyalty.Setting
		qs         = ormer.QueryTable(new(loyalty.Setting))
		fromRecord int32
	)
	fromRecord = pageIndex*pageSize - pageSize
	qs = qs.Limit(pageSize, fromRecord)
	if _, err := qs.All(&list); err != nil {
		return nil, err
	}

	return list, nil
}
