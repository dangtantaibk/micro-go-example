package controllers

import (
	"tng/common/logger"
	"tng/menu-service/dtos"
	"tng/menu-service/services"
)

type AreaController struct {
	BaseController
	areaService services.AreaService
}

func (a *AreaController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.AreaService) {
		a.areaService = s
	})
}

// Get List Area
// @Title Get List Area.
// @Description Get List Area
// @Param merchant_code	query string true "Merchant code"
// @Success 200 {object} dtos.ListAreaResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /list [get]
func (a *AreaController) List() {
	var request dtos.ListAreaRequest
	if err := a.ParseForm(&request); err != nil {
		logger.Errorf(a.Ctx.Request.Context(), "Parsing form: %v", err)
		a.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	if request.MerchantCode == "" {
		logger.Errorf(a.Ctx.Request.Context(), "MerchantCode invalid")
		a.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	a.Respond(a.areaService.List(a.Ctx.Request.Context(), &request))
}
