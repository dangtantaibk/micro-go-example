
package repositories

import (
	"context"
	"github.com/astaxie/beego/orm"
	"tng/common/models/loyalty"
	"tng/common/utils/db"
)

type ClassTrackingRepository interface {
	InsertOrUpdate(context.Context, *db.DB, *loyalty.ClassTracking) error
	Delete(context.Context, *db.DB, int64) error
	GetByID(context.Context, *db.DB, int64) (*loyalty.ClassTracking, error)
	List(context.Context, *db.DB, int32, int32) ([]*loyalty.ClassTracking, error)
}

type classTrackingRepository struct{}

func NewClassTrackingRepository(dbFactory db.Factory) ClassTrackingRepository {
	return &classTrackingRepository{}
}

func (r *classTrackingRepository) InsertOrUpdate(ctx context.Context, ormer *db.DB, dataInput *loyalty.ClassTracking) error {
	_, err := ormer.InsertOrUpdate(dataInput)
	return err
}

func (r *classTrackingRepository) Delete(ctx context.Context, ormer *db.DB, id int64) error {
	qs := ormer.QueryTable(new(loyalty.ClassTracking))
	cond := orm.NewCondition().
		And("id", id)
	_, err := qs.SetCond(cond).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r *classTrackingRepository) GetByID(ctx context.Context, ormer *db.DB, id int64) (*loyalty.ClassTracking, error) {
	cond := orm.NewCondition().
		And("id", id)
	item := &loyalty.ClassTracking{}
	qs := ormer.QueryTable(new(loyalty.ClassTracking))
	err := qs.SetCond(cond).One(item)
	if err != nil {
		return nil, err
	}
	return item, nil

}

func (r *classTrackingRepository) List(ctx context.Context, ormer *db.DB, pageIndex int32, pageSize int32) ([]*loyalty.ClassTracking, error) {
	var (
		list       []*loyalty.ClassTracking                
		qs         = ormer.QueryTable(new(loyalty.ClassTracking))
		fromRecord int32
	)
	fromRecord = pageIndex*pageSize - pageSize
	qs = qs.Limit(pageSize, fromRecord)
	if _, err := qs.All(&list); err != nil {
		return nil, err
	}

	return list, nil
}
