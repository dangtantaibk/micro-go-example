
package controllers

import (
	"encoding/json"
	"tng/common/logger"
	"tng/loyalty-service/dtos"
	"tng/loyalty-service/services"
)

type PointClassController struct {
	BaseController
	pointClassService services.PointClassService
}

func (c *PointClassController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.PointClassService) {
		c.pointClassService = s
	})
}

// InsertOrUpdate pointClass.
// @Title InsertOrUpdate pointClass// @Description InsertOrUpdate pointClass for loyalty
// @Param	body	body	dtos.InsertOrUpdatePointClassRequest	true	"Information PointClass"
// @Success 200 {object} dtos.InsertOrUpdatePointClassResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /insertorupdate [post]
func (c *PointClassController) InsertOrUpdate(body *dtos.InsertOrUpdatePointClassRequest) {
	//if body.ID <= 0 {
	//	logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
	//	c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
	//	return
	//}
	
	c.Respond(c.pointClassService.InsertOrUpdate(c.Ctx.Request.Context(), body))
}

// Delete pointClass.
// @Title Delete pointClass// @Description Delete pointClass for loyalty
// @Param	body	body	dtos.DeletePointClassRequest	true	"Information PointClass"
// @Success 200 {object} dtos.DeletePointClassResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /delete [post]
func (c *PointClassController) DeletePointClass(body *dtos.DeletePointClassRequest) {
	//if body.ID <= 0 {
	//	logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
	//	c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
	//	return
	//
	c.Respond(c.pointClassService.Delete(c.Ctx.Request.Context(), body))
}

// GetByID pointClass.
// @Title GetByID pointClass// @Description GetByID pointClass loyalty
// @Param id		query	int		true	"ID pointClass"
// @Success 200 {object} dtos.GetPointClassResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /getbyid [get]
func (c *PointClassController) GetByID() {
	var req dtos.GetPointClassRequest
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

	c.Respond(c.pointClassService.GetByID(c.Ctx.Request.Context(), &req))
}

// List pointClass.
// @Title List pointClass// @Description List pointClass loyalty
// @Param	body	body	dtos.MetaRequest	true	"data={\"page_index\":1,\"page_size\":100}"
// @Success 200 {object} dtos.ListPointClassResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /list [post]
func (c *PointClassController) ListPointClass(body *dtos.MetaRequest) {
	var request dtos.ListPointClassRequest
	if body.Data == "" {
		logger.Errorf(c.Ctx.Request.Context(), "Data invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	err := json.Unmarshal([]byte(body.Data), &request)
	if err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Unmarshal data error")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	// validate request neu can thiet
	// do something
	//--------------------------------
	/*if !c.CheckMeta(body) {
		logger.Errorf(c.Ctx.Request.Context(), "Request Invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}*/

	if request.PageIndex <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "PageIndex invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	if request.PageSize <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "PageSize invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	c.Respond(c.pointClassService.List(c.Ctx.Request.Context(), &request))
}

