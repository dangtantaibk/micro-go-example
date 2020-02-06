
package repositories

import (
	"context"
	"github.com/astaxie/beego/orm"
	"tng/common/models/loyalty"
	"tng/common/utils/db"
)

type PointTypeRepository interface {
	InsertOrUpdate(context.Context, *db.DB, *loyalty.PointType) error
	Delete(context.Context, *db.DB, string) error
	GetByID(context.Context, *db.DB, string) (*loyalty.PointType, error)
	List(context.Context, *db.DB, int32, int32) ([]*loyalty.PointType, error)
}

type pointTypeRepository struct{}

func NewPointTypeRepository(dbFactory db.Factory) PointTypeRepository {
	return &pointTypeRepository{}
}

func (r *pointTypeRepository) InsertOrUpdate(ctx context.Context, ormer *db.DB, dataInput *loyalty.PointType) error {
	_, err := ormer.InsertOrUpdate(dataInput)
	return err
}

func (r *pointTypeRepository) Delete(ctx context.Context, ormer *db.DB, id string) error {
	qs := ormer.QueryTable(new(loyalty.PointType))
	cond := orm.NewCondition().
		And("id", id)
	_, err := qs.SetCond(cond).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r *pointTypeRepository) GetByID(ctx context.Context, ormer *db.DB, id string) (*loyalty.PointType, error) {
	cond := orm.NewCondition().
		And("id", id)
	item := &loyalty.PointType{}
	qs := ormer.QueryTable(new(loyalty.PointType))
	err := qs.SetCond(cond).One(item)
	if err != nil {
		return nil, err
	}
	return item, nil

}

func (r *pointTypeRepository) List(ctx context.Context, ormer *db.DB, pageIndex int32, pageSize int32) ([]*loyalty.PointType, error) {
	var (
		list       []*loyalty.PointType                
		qs         = ormer.QueryTable(new(loyalty.PointType))
		fromRecord int32
	)
	fromRecord = pageIndex*pageSize - pageSize
	qs = qs.Limit(pageSize, fromRecord)
	if _, err := qs.All(&list); err != nil {
		return nil, err
	}

	return list, nil
}
