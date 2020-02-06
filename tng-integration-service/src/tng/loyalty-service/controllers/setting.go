package controllers

import (
	"encoding/json"
	"fmt"
	"tng/common/logger"
	"tng/loyalty-service/dtos"
	"tng/loyalty-service/services"
)

type SettingController struct {
	BaseController
	settingService services.SettingService
}

func (c *SettingController) Prepare() {
	fmt.Println("Prepare")
	//c.BaseController.Prepare()
	_ = services.GetServiceContainer().Invoke(func(s services.SettingService) {
		c.settingService = s
	})
}

// InsertOrUpdate setting.
// @Title InsertOrUpdate setting
// @Description InsertOrUpdate setting for loyalty
// @Param	body	body	dtos.MetaRequest	true	"Information Setting"
// @Success 200 {object} dtos.InsertOrUpdateSettingResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /insertorupdate [post]
func (c *SettingController) InsertOrUpdate(body *dtos.MetaRequest) {
	request := &dtos.InsertOrUpdateSettingRequest{}
	if body.Data == "" {
		logger.Errorf(c.Ctx.Request.Context(), "Data invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	err := json.Unmarshal([]byte(body.Data), request)
	if err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Unmarshal data error")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	// validate request neu can thiet
	// do something
	//--------------------------------
	if !c.CheckMeta(body) {
		logger.Errorf(c.Ctx.Request.Context(), "Request Invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	c.Respond(c.settingService.InsertOrUpdate(c.Ctx.Request.Context(), request))
}

// Delete setting.
// @Title Delete setting
// @Description Delete setting for loyalty
// @Param	body	body	dtos.DeleteSettingRequest	true	"Information Setting"
// @Success 200 {object} dtos.DeleteSettingResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /delete [post]
func (c *SettingController) DeleteSetting(body *dtos.DeleteSettingRequest) {
	if body.ID <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	c.Respond(c.settingService.Delete(c.Ctx.Request.Context(), body))
}

// GetByID setting.
// @Title GetByID setting
// @Description GetByID setting loyalty
// @Param mc			query	string		true	"Merchant Code"
// @Param ts			query	int			true	"TimeStamp"
// @Param user_agent	query	string		true	"User Agent"
// @Param client_ip		query	string		true	"Client IP"
// @Param check_sum		query	string		true	"Checksum"
// @Param id			query	int			true	"ID Setting"
// @Success 200 {object} dtos.GetSettingResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /getbyid [get]
func (c *SettingController) GetByID() {
	var req dtos.GetSettingRequest
	if err := c.ParseForm(&req); err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Parsing form: %v", err)
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	if req.ID <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	c.Respond(c.settingService.GetByID(c.Ctx.Request.Context(), &req))
}

// List setting.
// @Title List setting
// @Description List setting loyalty
// @Param page_index	query	string	true	"Page Index"
// @Param page_size		query	int		true	"Page Size"
// @Success 200 {object} dtos.ListSettingResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /list [get]
func (c *SettingController) ListSetting() {
	var req dtos.ListSettingRequest
	if err := c.ParseForm(&req); err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Parsing form: %v", err)
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	if req.PageIndex <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "PageIndex invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	if req.PageSize <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "PageSize invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	c.Respond(c.settingService.List(c.Ctx.Request.Context(), &req))
}
