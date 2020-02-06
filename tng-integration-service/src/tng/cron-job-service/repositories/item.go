package repositories

import (
	"context"
	"github.com/astaxie/beego/orm"
	cron_job "tng/common/models/cron-job"
	"tng/common/utils/db"
)

type ItemRepository interface {
	GetDetails(context.Context, *db.DB, int) (*cron_job.Item, error)
}

type itemRepository struct{}

func NewItemRepository() ItemRepository {
	return &itemRepository{}
}

func (i itemRepository) GetDetails(ctx context.Context, ormer *db.DB, itemId int) (*cron_job.Item, error) {
	var (
		item cron_job.Item
		qs   = ormer.QueryTable(new(cron_job.Item))
		cond = orm.NewCondition().
			And("auto_sync", 1).
			And("item_id", itemId)
	)
	if err := qs.SetCond(cond).One(&item); err != nil {
		return nil, err
	}

	return &item, nil
}
