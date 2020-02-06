package controllers

import (
	"encoding/json"
	"tng/common/logger"
	"tng/loyalty-service/dtos"
	"tng/loyalty-service/services"
)

type PointController struct {
	BaseController
	pointService services.PointService
}

func (c *PointController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.PointService) {
		c.pointService = s
	})
}

// InsertOrUpdate point.
// @Title InsertOrUpdate point// @Description InsertOrUpdate point for loyalty
// @Param	body	body	dtos.InsertOrUpdatePointRequest	true	"Information Point"
// @Success 200 {object} dtos.InsertOrUpdatePointResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /insertorupdate [post]
func (c *PointController) InsertOrUpdate(body *dtos.InsertOrUpdatePointRequest) {
	//if body.ID <= 0 {
	//	logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
	//	c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
	//	return
	//}

	c.Respond(c.pointService.InsertOrUpdate(c.Ctx.Request.Context(), body))
}

// Delete point.
// @Title Delete point// @Description Delete point for loyalty
// @Param	body	body	dtos.DeletePointRequest	true	"Information Point"
// @Success 200 {object} dtos.DeletePointResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /delete [post]
func (c *PointController) DeletePoint(body *dtos.DeletePointRequest) {
	//if body.ID <= 0 {
	//	logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
	//	c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
	//	return
	//
	c.Respond(c.pointService.Delete(c.Ctx.Request.Context(), body))
}

// GetByID point.
// @Title GetByID point// @Description GetByID point loyalty
// @Param id		query	int		true	"ID point"
// @Success 200 {object} dtos.GetPointResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /getbyid [get]
func (c *PointController) GetByID() {
	var req dtos.GetPointRequest
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

	c.Respond(c.pointService.GetByID(c.Ctx.Request.Context(), &req))
}

// List point.
// @Title List point// @Description List point loyalty
// @Param page_index	query	string	true	"Page Index"
// @Param page_size		query	int		true	"Page Size"
// @Success 200 {object} dtos.ListPointResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /list [get]
func (c *PointController) ListPoint() {
	var req dtos.ListPointRequest
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

	c.Respond(c.pointService.List(c.Ctx.Request.Context(), &req))
}

// AddPoint point.
// @Title AddPoint point// @Description AddPoint point for loyalty
// @Param	body	body	dtos.MetaRequest	true	"data={"app_id": "string","campaign_code": "string","channel": "string","for_transaction_id": "string","json_detail": "string","notes": "string","promotion_percent": 0,"source": "string","transaction_amount": 0,"user_id": "string"}"
// @Success 200 {object} dtos.AddPointResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /addpoint [post]
func (c *PointController) AddPoint(body *dtos.MetaRequest) {
	request := &dtos.AddPointRequest{}
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
	/*if !c.CheckMeta(body) {
		logger.Errorf(c.Ctx.Request.Context(), "Request Invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}*/
	c.Respond(c.pointService.AddPoint(c.Ctx.Request.Context(), request))
}

// Search point.
// @Title Search point
// @Description Search point loyalty
// @Param term	query	string	false	"Term"
// @Param point	query	int64	false	"Point"
// @Param transaction_amount	query	int64	false	"Transaction Amount"
// @Param created_from	query	string	false	"Created From"
// @Param created_to	query	string	false	"Created To"
// @Param status	query	int32	false	"Status"
// @Param promotion_percent	query	float64	false	"Promotion Percent"
// @Param rate	query	float64	false	"Rate"
// @Param sort_column	query	string	false	"Sort Column"
// @Param sort_direction	query	string	false	"Sort Direction"
// @Param page_index	query	int	true	"Page Index"
// @Param page_size		query	int		true	"Page Size"
// @Success 200 {object} dtos.SearchPointResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /search [get]
func (c *PointController) SearchPoint() {
	var req dtos.SearchPointRequest
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

	c.Respond(c.pointService.Search(c.Ctx.Request.Context(), &req))
}

// Check Point.
// @Title CheckPoint point// @Description CheckPoint point for loyalty
// @Param	body	body	dtos.MetaRequest	true	"data=user_id"
// @Success 200 {object} dtos.CheckPointResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /checkpoint [post]
func (c *PointController) CheckPoint(body *dtos.MetaRequest) {
	request := &dtos.CheckPointRequest{}
	if body.Data == "" {
		logger.Errorf(c.Ctx.Request.Context(), "Data invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	//request.UserId = body.Data;
	err := json.Unmarshal([]byte(body.Data), request)
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

	c.Respond(c.pointService.CheckPoint(c.Ctx.Request.Context(), request))
}

// PointHistory.
// @Title PointHistory point// @Description PointHistory point loyalty
// @Param	body	body	dtos.MetaRequest	true	"data={\"page_index\":1,\"page_size\":100,\"user_id\":\"quannd3\",}"
// @Success 200 {object} dtos.PointHistoryResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /pointhistory [post]
func (c *PointController) PointHistory(body *dtos.MetaRequest) {
	var request dtos.PointHistoryRequest
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

	if request.UserId == "" {
		logger.Errorf(c.Ctx.Request.Context(), "UserId invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	c.Respond(c.pointService.PointHistory(c.Ctx.Request.Context(), &request))
}

// AddPoint point.
// @Title CheckOldPoint point// @Description CheckOldPoint point for loyalty
// @Param	body	body	dtos.MetaRequest	true	"data={\"user_id\":\"quannd3\",\"for_transaction_id\":\"111\",\"app_id\":\"111\",}"
// @Success 200 {object} dtos.CheckOldPointResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /checkoldpoint [post]
func (c *PointController) CheckOldPoint(body *dtos.MetaRequest) {
	request := &dtos.CheckOldPointRequest{}
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
	/*if !c.CheckMeta(body) {
		logger.Errorf(c.Ctx.Request.Context(), "Request Invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}*/

	c.Respond(c.pointService.CheckOldPoint(c.Ctx.Request.Context(), request))
}
