
package controllers

import (
	"tng/common/logger"
	"tng/loyalty-service/dtos"
	"tng/loyalty-service/services"
)

type ClassTrackingController struct {
	BaseController
	classTrackingService services.ClassTrackingService
}

func (c *ClassTrackingController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.ClassTrackingService) {
		c.classTrackingService = s
	})
}

// InsertOrUpdate classTracking.
// @Title InsertOrUpdate classTracking// @Description InsertOrUpdate classTracking for loyalty
// @Param	body	body	dtos.InsertOrUpdateClassTrackingRequest	true	"Information ClassTracking"
// @Success 200 {object} dtos.InsertOrUpdateClassTrackingResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /insertorupdate [post]
func (c *ClassTrackingController) InsertOrUpdate(body *dtos.InsertOrUpdateClassTrackingRequest) {
	//if body.ID <= 0 {
	//	logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
	//	c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
	//	return
	//}
	
	c.Respond(c.classTrackingService.InsertOrUpdate(c.Ctx.Request.Context(), body))
}

// Delete classTracking.
// @Title Delete classTracking// @Description Delete classTracking for loyalty
// @Param	body	body	dtos.DeleteClassTrackingRequest	true	"Information ClassTracking"
// @Success 200 {object} dtos.DeleteClassTrackingResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /delete [post]
func (c *ClassTrackingController) DeleteClassTracking(body *dtos.DeleteClassTrackingRequest) {
	//if body.ID <= 0 {
	//	logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
	//	c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
	//	return
	//
	c.Respond(c.classTrackingService.Delete(c.Ctx.Request.Context(), body))
}

// GetByID classTracking.
// @Title GetByID classTracking// @Description GetByID classTracking loyalty
// @Param id		query	int		true	"ID classTracking"
// @Success 200 {object} dtos.GetClassTrackingResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /getbyid [get]
func (c *ClassTrackingController) GetByID() {
	var req dtos.GetClassTrackingRequest
	if err := c.ParseForm(&req); err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Parsing form: %v", err)
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	//if req.ID <= 0 {
	//	logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
	//	c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
	//	return
	//

	c.Respond(c.classTrackingService.GetByID(c.Ctx.Request.Context(), &req))
}

// List classTracking.
// @Title List classTracking// @Description List classTracking loyalty
// @Param page_index	query	string	true	"Page Index"
// @Param page_size		query	int		true	"Page Size"
// @Success 200 {object} dtos.ListClassTrackingResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /list [get]
func (c *ClassTrackingController) ListClassTracking() {
	var req dtos.ListClassTrackingRequest
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

	c.Respond(c.classTrackingService.List(c.Ctx.Request.Context(), &req))
}

