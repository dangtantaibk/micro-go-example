
package repositories

import (
	"context"
	"github.com/astaxie/beego/orm"
	"tng/common/models/loyalty"
	"tng/common/utils/db"
)

type PointClassRepository interface {
	InsertOrUpdate(context.Context, *db.DB, *loyalty.PointClass) error
	Delete(context.Context, *db.DB, int64) error
	GetByID(context.Context, *db.DB, int64) (*loyalty.PointClass, error)
	List(context.Context, *db.DB, int32, int32) ([]*loyalty.PointClass, error)
	GetAll(*db.DB) ([]*loyalty.PointClass, error)
}

type pointClassRepository struct{}

func NewPointClassRepository(dbFactory db.Factory) PointClassRepository {
	return &pointClassRepository{}
}

func (r *pointClassRepository) InsertOrUpdate(ctx context.Context, ormer *db.DB, dataInput *loyalty.PointClass) error {
	_, err := ormer.InsertOrUpdate(dataInput)
	return err
}

func (r *pointClassRepository) Delete(ctx context.Context, ormer *db.DB, id int64) error {
	qs := ormer.QueryTable(new(loyalty.PointClass))
	cond := orm.NewCondition().
		And("id", id)
	_, err := qs.SetCond(cond).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r *pointClassRepository) GetByID(ctx context.Context, ormer *db.DB, id int64) (*loyalty.PointClass, error) {
	cond := orm.NewCondition().
		And("id", id)
	item := &loyalty.PointClass{}
	qs := ormer.QueryTable(new(loyalty.PointClass))
	err := qs.SetCond(cond).One(item)
	if err != nil {
		return nil, err
	}
	return item, nil

}

func (r *pointClassRepository) List(ctx context.Context, ormer *db.DB, pageIndex int32, pageSize int32) ([]*loyalty.PointClass, error) {
	var (
		list       []*loyalty.PointClass                
		qs         = ormer.QueryTable(new(loyalty.PointClass))
		fromRecord int32
	)
	fromRecord = pageIndex*pageSize - pageSize
	qs = qs.Limit(pageSize, fromRecord)
	if _, err := qs.All(&list); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *pointClassRepository) GetAll(ormer *db.DB) ([]*loyalty.PointClass, error) {
	var (
		list       []*loyalty.PointClass
		qs         = ormer.QueryTable(new(loyalty.PointClass))
	)
	if _, err := qs.All(&list); err != nil {
		return nil, err
	}

	return list, nil
}