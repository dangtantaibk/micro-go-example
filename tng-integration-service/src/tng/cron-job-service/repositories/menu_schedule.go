package repositories

import (
	"context"
	"github.com/astaxie/beego/orm"
	"tng/common/location"
	cron_job "tng/common/models/cron-job"
	"tng/common/utils/db"
)

type ScheduleRepository interface {
	GetList(context.Context, *db.DB) ([]*cron_job.Schedule, error)
	Insert(context.Context, *db.DB, *cron_job.Schedule) error
	Delete(context.Context, *db.DB) error
}

type scheduleRepository struct{}

func NewScheduleRepository() ScheduleRepository {
	return &scheduleRepository{}
}

func (r *scheduleRepository) Delete(ctx context.Context, ormer *db.DB) error {
	t := location.GetVNCurrentTime()
	nowDay := t.Day()
	_, err := ormer.QueryTable(new(cron_job.Schedule)).
		Filter("day_of_schedule", nowDay).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r *scheduleRepository) GetList(ctx context.Context, ormer *db.DB) ([]*cron_job.Schedule, error) {
	t := location.GetVNCurrentTimeYMD(0, 0, -1)
	yesterday := t.Day()
	var (
		list []*cron_job.Schedule
		qs   = ormer.QueryTable(new(cron_job.Schedule))
		cond = orm.NewCondition().And("day_of_schedule", yesterday)
	)
	if _, err := qs.SetCond(cond).All(&list); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *scheduleRepository) Insert(ctx context.Context, ormer *db.DB, schedule *cron_job.Schedule) error {
	_, err := ormer.InsertOrUpdate(schedule)
	return err
}
