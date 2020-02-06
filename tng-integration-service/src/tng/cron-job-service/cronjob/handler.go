package cronjob

import (
	"context"
	"github.com/robfig/cron"
	"tng/common/location"
	"tng/common/logger"
)

type cronJobHandler struct {
	cron *cron.Cron
	job  ExcuteJob
}

// CronHandler init handler
type CronHandler interface {
	InitCronJob(context.Context)
	Run()
}

// NewCronHandler init hdl
func NewCronHandler(ctx context.Context) CronHandler {
	opts := cron.WithLocation(location.VNLocation)
	h := &cronJobHandler{
		cron: cron.New(opts),
		job:  NewExcuteJobHandler(),
	}
	return h
}

func (h *cronJobHandler) InitCronJob(ctx context.Context) {
	_, err := h.cron.AddFunc(ScheduleDayByDayMenuSchedule, func() {
		h.job.SyncMenuSchedule()
	})
	if err != nil {
		logger.Errorf(ctx, "InitCronJob ScheduleDayByDayMenuSchedule error: %v", err)
	}
	_, err = h.cron.AddFunc(ScheduleDayByDayUpdateMenuSchedule, func() {
		h.job.NotifyUpdateMenuSchedule()
	})
	if err != nil {
		logger.Errorf(ctx, "InitCronJob ScheduleDayByDayUpdateMenuSchedule error: %v", err)
	}
}

func (h *cronJobHandler) Run() {
	h.cron.Run()
}
