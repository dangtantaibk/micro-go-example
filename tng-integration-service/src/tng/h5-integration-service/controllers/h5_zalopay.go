package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"tng/common/logger"
	"tng/common/utils/cfgutil"
	"tng/common/utils/hashutil"
	"tng/common/utils/keysutil"
	"tng/h5-integration-service/dtos"
	"tng/h5-integration-service/services"
)

var (
	runMode = cfgutil.Load("RunMode")
)

// H5ZaloPayController represents controller of Token
type H5ZaloPayController struct {
	BaseController
	h5ZaloPayService services.H5ZaloPayService
}

// Prepare handles prepare of H5ZaloPayController.
func (c *H5ZaloPayController) Prepare() {
	//c.BaseController.Prepare()
	_ = services.GetServiceContainer().Invoke(func(s services.H5ZaloPayService) {
		c.h5ZaloPayService = s
	})
}

// GrantedMBToken a MBToken from H5ZaloPay.
// @Title GrantedMBToken a MBToken from H5ZaloPay
// @Description GrantedMBToken a MBToken from H5ZaloPay or cache
// @Param user_id	query	string	true	"UserID for getting."
// @Param app_id	query	int		true	"AppID for getting."
// @Param ma_token	query	string	true	"MAToken for getting."
// @Param timestamp	query	int		true	"TimeStamp for current request. Ex: 1572849166"
// @Param sig		query	string	true	"Sig for SHA256. input: AppId + "|" + ClientKey + "|" + MAToken + "|" + TimeStamp"
// @Success 200 {object} dtos.GetH5ZaloPayResp
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /grantedmbtoken [get]
func (c *H5ZaloPayController) GrantedMBToken() {
	var req dtos.GetH5ZaloPayRequest
	if err := c.ParseForm(&req); err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Parsing form: %v", err)
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	if req.AppID <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "AppID invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if req.UserID == "" {
		logger.Errorf(c.Ctx.Request.Context(), "UserID invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if req.MAToken == "" {
		logger.Errorf(c.Ctx.Request.Context(), "MAToken invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if req.TimeStamp <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "TimeStamp invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if req.Sig == "" {
		logger.Errorf(c.Ctx.Request.Context(), "Sig invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	clientKey := keysutil.GetKey(fmt.Sprintf("%d", req.AppID))
	dataInput := fmt.Sprintf("%d|%s|%s|%d", req.AppID, clientKey, req.MAToken, req.TimeStamp)
	sig := hashutil.GetSHA256(dataInput)
	if sig != req.Sig && runMode != beego.DEV {
		logger.Errorf(c.Ctx.Request.Context(), "Sig not matching")
		c.Respond(nil, dtos.NewAppError(dtos.ErrorSigNotMatching))
		return
	}

	c.Respond(c.h5ZaloPayService.GrantedMBToken(c.Ctx.Request.Context(), &req))
}

// GetPaymentOrderUrl from H5ZaloPay.
// @Title GetPaymentOrderUrl from H5ZaloPay
// @Description Get PaymentUrl from H5ZaloPay, Sig for SHA256. input: AppID + "|" + clientKey + "|" + UserID + "|" + TimeStamp
// @Param	body	body	dtos.H5ZaloPayOrderURLReq	true	"Information Order URL"
// @Success 200 {object} dtos.H5ZaloPayOrderURLResp
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /paymenturl [post]
func (c *H5ZaloPayController) GetPaymentOrderUrl(body *dtos.H5ZaloPayOrderURLReq) {
	if body.UserID == "" {
		logger.Errorf(c.Ctx.Request.Context(), "UserID invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if body.ZPOrderURL == "" {
		logger.Errorf(c.Ctx.Request.Context(), "ZPOrderURL invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if body.AppID <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "AppID invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if body.MAToken == "" {
		logger.Errorf(c.Ctx.Request.Context(), "ZPOrderURL invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if body.TimeStamp <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "MAToken invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if body.Sig == "" {
		logger.Errorf(c.Ctx.Request.Context(), "Sig invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	clientKey := keysutil.GetKey(fmt.Sprintf("%d", body.AppID))
	dataInput := fmt.Sprintf("%d|%s|%s|%d", body.AppID, clientKey, body.UserID, body.TimeStamp)
	sig := hashutil.GetSHA256(dataInput)
	if sig != body.Sig && runMode != beego.DEV {
		logger.Errorf(c.Ctx.Request.Context(), "Sig not matching")
		c.Respond(nil, dtos.NewAppError(dtos.ErrorSigNotMatching))
		return
	}

	c.Respond(c.h5ZaloPayService.GetPaymentOrderUrl(c.Ctx.Request.Context(), body))
}
