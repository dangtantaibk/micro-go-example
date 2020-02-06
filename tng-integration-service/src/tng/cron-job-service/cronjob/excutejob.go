package cronjob

import (
	"context"
	"tng/cron-job-service/services"
)

type ExcuteJob interface {
	SyncMenuSchedule()
	NotifyUpdateMenuSchedule()
}

type excuteJobHandler struct {
	scheduleService services.ScheduleService
}

// NewExcuteJobHandler int job
func NewExcuteJobHandler() ExcuteJob {
	h := &excuteJobHandler{}
	_ = services.GetServiceContainer().Invoke(func(s services.ScheduleService) {
		h.scheduleService = s
	})
	return h
}

func (h *excuteJobHandler) SyncMenuSchedule() {
	h.scheduleService.Sync(context.Background())
}

func (h *excuteJobHandler) NotifyUpdateMenuSchedule() {
	h.scheduleService.NotifyUpdateMenu(context.Background())
}