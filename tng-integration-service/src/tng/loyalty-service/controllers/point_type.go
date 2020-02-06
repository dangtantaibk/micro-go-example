
package controllers

import (
	"tng/common/logger"
	"tng/loyalty-service/dtos"
	"tng/loyalty-service/services"
)

type PointTypeController struct {
	BaseController
	pointTypeService services.PointTypeService
}

func (c *PointTypeController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.PointTypeService) {
		c.pointTypeService = s
	})
}

// InsertOrUpdate pointType.
// @Title InsertOrUpdate pointType// @Description InsertOrUpdate pointType for loyalty
// @Param	body	body	dtos.InsertOrUpdatePointTypeRequest	true	"Information PointType"
// @Success 200 {object} dtos.InsertOrUpdatePointTypeResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /insertorupdate [post]
func (c *PointTypeController) InsertOrUpdate(body *dtos.InsertOrUpdatePointTypeRequest) {
	//if body.ID <= 0 {
	//	logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
	//	c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
	//	return
	//}
	
	c.Respond(c.pointTypeService.InsertOrUpdate(c.Ctx.Request.Context(), body))
}

// Delete pointType.
// @Title Delete pointType// @Description Delete pointType for loyalty
// @Param	body	body	dtos.DeletePointTypeRequest	true	"Information PointType"
// @Success 200 {object} dtos.DeletePointTypeResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /delete [post]
func (c *PointTypeController) DeletePointType(body *dtos.DeletePointTypeRequest) {
	//if body.ID <= 0 {
	//	logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
	//	c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
	//	return
	//
	c.Respond(c.pointTypeService.Delete(c.Ctx.Request.Context(), body))
}

// GetByID pointType.
// @Title GetByID pointType// @Description GetByID pointType loyalty
// @Param id		query	int		true	"ID pointType"
// @Success 200 {object} dtos.GetPointTypeResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /getbyid [get]
func (c *PointTypeController) GetByID() {
	var req dtos.GetPointTypeRequest
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

	c.Respond(c.pointTypeService.GetByID(c.Ctx.Request.Context(), &req))
}

// List pointType.
// @Title List pointType// @Description List pointType loyalty
// @Param page_index	query	string	true	"Page Index"
// @Param page_size		query	int		true	"Page Size"
// @Success 200 {object} dtos.ListPointTypeResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /list [get]
func (c *PointTypeController) ListPointType() {
	var req dtos.ListPointTypeRequest
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

	c.Respond(c.pointTypeService.List(c.Ctx.Request.Context(), &req))
}

