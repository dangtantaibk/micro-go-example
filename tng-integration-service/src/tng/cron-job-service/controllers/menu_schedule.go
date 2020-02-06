package controllers

import (
	"tng/cron-job-service/dtos"
	"tng/cron-job-service/services"
)

var (
//runMode = cfgutil.Load("RUNMODE")
)

// WarmUpScheduleController represents controller of Token
type ScheduleController struct {
	BaseController
	scheduleService services.ScheduleService
}

// Prepare handles prepare of MenuScheduleController.
func (c *ScheduleController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.ScheduleService) {
		c.scheduleService = s
	})
}

// WarmUp Schedule.
// @Title WarmUp Schedule
// @Description Manual WarmUpSchedule
// @Param	body	body	dtos.WarmUpScheduleRequest	true	"if merchant_code is_empty -> get all"
// @Success 200 {object} dtos.WarmUpScheduleResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /warm-up [post]
func (c *ScheduleController) WarmUp(body *dtos.WarmUpScheduleRequest) {
	c.Respond(c.scheduleService.WarmUp(c.Ctx.Request.Context(), body))
}

