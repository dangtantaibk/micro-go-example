
package controllers

import (
	"tng/common/logger"
	"tng/loyalty-service/dtos"
	"tng/loyalty-service/services"
)

type WalletController struct {
	BaseController
	walletService services.WalletService
}

func (c *WalletController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.WalletService) {
		c.walletService = s
	})
}

// InsertOrUpdate wallet.
// @Title InsertOrUpdate wallet// @Description InsertOrUpdate wallet for loyalty
// @Param	body	body	dtos.InsertOrUpdateWalletRequest	true	"Information Wallet"
// @Success 200 {object} dtos.InsertOrUpdateWalletResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /insertorupdate [post]
func (c *WalletController) InsertOrUpdate(body *dtos.InsertOrUpdateWalletRequest) {
	//if body.ID <= 0 {
	//	logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
	//	c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
	//	return
	//}
	
	c.Respond(c.walletService.InsertOrUpdate(c.Ctx.Request.Context(), body))
}

// Delete wallet.
// @Title Delete wallet// @Description Delete wallet for loyalty
// @Param	body	body	dtos.DeleteWalletRequest	true	"Information Wallet"
// @Success 200 {object} dtos.DeleteWalletResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /delete [post]
func (c *WalletController) DeleteWallet(body *dtos.DeleteWalletRequest) {
	//if body.ID <= 0 {
	//	logger.Errorf(c.Ctx.Request.Context(), "ID invalid")
	//	c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
	//	return
	//
	c.Respond(c.walletService.Delete(c.Ctx.Request.Context(), body))
}

// GetByID wallet.
// @Title GetByID wallet// @Description GetByID wallet loyalty
// @Param id		query	int		true	"ID wallet"
// @Success 200 {object} dtos.GetWalletResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /getbyid [get]
func (c *WalletController) GetByID() {
	var req dtos.GetWalletRequest
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

	c.Respond(c.walletService.GetByID(c.Ctx.Request.Context(), &req))
}

// List wallet.
// @Title List wallet// @Description List wallet loyalty
// @Param page_index	query	string	true	"Page Index"
// @Param page_size		query	int		true	"Page Size"
// @Success 200 {object} dtos.ListWalletResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /list [get]
func (c *WalletController) ListWallet() {
	var req dtos.ListWalletRequest
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

	c.Respond(c.walletService.List(c.Ctx.Request.Context(), &req))
}

