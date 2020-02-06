package controllers

import (
	"tng/common/logger"
	"tng/menu-service/dtos"
	"tng/menu-service/services"
)

type ItemTypeController struct {
	BaseController
	itemTypeService services.ItemTypeService
}

func (a *ItemTypeController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.ItemTypeService) {
		a.itemTypeService = s
	})
}

// Get List Item Type
// @Title Get List Item Type.
// @Description Get List Item Type
// @Param merchant_code	query string true "Merchant code"
// @Success 200 {object} dtos.ListItemTypeResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /list [get]
func (i *ItemTypeController) List() {
	var request dtos.ListItemTypeRequest
	if err := i.ParseForm(&request); err != nil {
		logger.Errorf(i.Ctx.Request.Context(), "Parsing form: %v", err)
		i.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.MerchantCode == "" {
		logger.Errorf(i.Ctx.Request.Context(), "MerchantCode invalid")
		i.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	i.Respond(i.itemTypeService.List(i.Ctx.Request.Context(), &request))
}
// Create Item Type.
// @Title Create Item Type
// @Description Create Item Type
// @Param	body	body	dtos.CreateItemTypeRequest	true	"Item Type info"
// @Success 200 {object} dtos.CreateItemTypeResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /create [post]
func (i *ItemTypeController) Create(body *dtos.CreateItemTypeRequest) {
	if body.MerchantCode == "" {
		logger.Errorf(i.Ctx.Request.Context(), "MerchantCode invalid")
		i.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	i.Respond(i.itemTypeService.Create(i.Ctx.Request.Context(), body))
}

// Delete Item Type
// @Title Delete Item Type
// @Description Delete Item Type
// @Param	body	body dtos.DeleteItemTypeRequest	true	"Update request info"
// @Success 200 {object} dtos.DeleteItemTypeResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /delete [post]
func (i *ItemTypeController) DeleteItemType(body *dtos.DeleteItemTypeRequest) {
	if body.MerchantCode == "" {
		logger.Errorf(i.Ctx.Request.Context(), "MerchantCode invalid")
		i.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	i.Respond(i.itemTypeService.Delete(i.Ctx.Request.Context(), body))
}

// Update Item Type
// @Title Update Item Type
// @Description Update Item Type
// @Param	body	body dtos.UpdateItemTypeRequest	true	"Update request info"
// @Success 200 {object} dtos.UpdateItemTypeResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /update [post]
func (i *ItemTypeController) Update(body *dtos.UpdateItemTypeRequest) {
	if body.MerchantCode == "" {
		logger.Errorf(i.Ctx.Request.Context(), "MerchantCode invalid")
		i.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	i.Respond(i.itemTypeService.Update(i.Ctx.Request.Context(), body))
}