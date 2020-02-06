package controllers

import (
	"tng/common/logger"
	"tng/shipper-service/dtos"
	"tng/shipper-service/services"
)

type ShipperController struct {
	BaseController
	shipperService services.ShipperService
}

func (u *ShipperController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.ShipperService) {
		u.shipperService = s
	})
}

// Sign up
// @Title Sign up for shipper
// @Description Register username and password
// @Param body body	dtos.SignUpShipperRequest	true	"Information request"
// @Success 200 {object} dtos.SignUpShipperResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /signup [post]
func (s *ShipperController) Signup(body *dtos.SignUpShipperRequest) {
	if body.MerchantCode == "" {
		logger.Errorf(s.Ctx.Request.Context(), "MerchantCode invalid")
		s.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	s.Respond(s.shipperService.SignUp(s.Ctx.Request.Context(), body))
}

// Get List Shipper Account.
// @Title Get List Shipper Account.
// @Description Get List Shipper Account
// @Param merchant_code	query	string	true	"Merchant code"
// @Param current_page	query	int32	false	"Current page"
// @Param total_trans_per_page	query	int32	false	"Total per page"
// @Success 200 {object} dtos.ListShipperResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /list [get]
func (s *ShipperController) List() {
	var request dtos.ListShipperRequest
	if err := s.ParseForm(&request); err != nil {
		logger.Errorf(s.Ctx.Request.Context(), "Parsing form: %v", err)
		s.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.MerchantCode == "" {
		logger.Errorf(s.Ctx.Request.Context(), "MerchantCode invalid")
		s.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	s.Respond(s.shipperService.List(s.Ctx.Request.Context(), &request))
}

// Delete Shipper Account
// @Title Delete Shipper Account
// @Description Delete Shipper Account
// @Param	body	body dtos.DeleteShipperAccountRequest	true	"Request Info"
// @Success 200 {object} dtos.DeleteShipperAccountResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /delete [post]
func (s *ShipperController) DeleteShipper(body *dtos.DeleteShipperAccountRequest) {
	if body.MerchantCode == "" {
		logger.Errorf(s.Ctx.Request.Context(), "MerchantCode invalid")
		s.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	s.Respond(s.shipperService.Delete(s.Ctx.Request.Context(), body))
}

// Update Shipper Account
// @Title Update Shipper Account
// @Description Update Shipper Account
// @Param	body	body dtos.UpdateShipperRequest	true	"Update request info"
// @Success 200 {object} dtos.UpdateShipperResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /update [post]
func (s *ShipperController) Update(body *dtos.UpdateShipperRequest) {
	if body.MerchantCode == "" {
		logger.Errorf(s.Ctx.Request.Context(), "MerchantCode invalid")
		s.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	s.Respond(s.shipperService.Update(s.Ctx.Request.Context(), body))
}
