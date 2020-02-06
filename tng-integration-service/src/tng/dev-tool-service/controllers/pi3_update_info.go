package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"tng/common/logger"
	"tng/common/utils/cfgutil"
	"tng/common/utils/hashutil"
	"tng/dev-tool-service/dtos"
	"tng/dev-tool-service/services"
)

var (
	runMode = cfgutil.Load("RunMode")
)

// Pi3UpdateInfoController represents controller of update info
type Pi3UpdateInfoController struct {
	BaseController
	pi3UpdateInfoService services.Pi3UpdateInfoService
}

// Prepare handles prepare of Pi3UpdateInfoController.
func (c *Pi3UpdateInfoController) Prepare() {
	c.BaseController.Prepare()
	_ = services.GetServiceContainer().Invoke(func(s services.Pi3UpdateInfoService) {
		c.pi3UpdateInfoService = s
	})
}

// GetPi3UpdateInfo from Dev Tool.
// @Title GetPi3UpdateInfo from Dev Tool
// @Description GetPi3UpdateInfo from Dev Tool
// @Param pos_id		query	int		true	"PosID for getting."
// @Param timestamp		query	int		true	"TimeStamp for current request. Ex: 1572849166"
// @Param sig			query	string	true	"Sig for SHA256. input: PosID + "|" + ClientKey + "|" + TimeStamp"
// @Success 200 {object} dtos.GetPi3UpdateInfoResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /get-pi3-update-info [get]
func (c *Pi3UpdateInfoController) GetPi3UpdateInfo() {
	var req dtos.GetPi3UpdateInfoRequest
	if err := c.ParseForm(&req); err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Parsing form: %v", err)
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	if req.PosID == "" {
		logger.Errorf(c.Ctx.Request.Context(), "PosID invalid")
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
	dataInput := fmt.Sprintf("%s|%s|%d", req.PosID, cfgutil.Load("TNG_CLIENT_KEY"), req.TimeStamp)
	sig := hashutil.GetSHA256(dataInput)
	if sig != req.Sig && runMode != beego.DEV {
		logger.Errorf(c.Ctx.Request.Context(), "Sig not matching")
		c.Respond(nil, dtos.NewAppError(dtos.ErrorSigNotMatching))
		return
	}

	c.Respond(c.pi3UpdateInfoService.GetPi3UpdateInfo(c.Ctx.Request.Context(), &req))
}

// RegisterPi3 from device pi3.
// @Title RegisterPi3 from device pi3
// @Description Register Device Pi3 when start device, Sig for SHA256. input: PosID + "|" + clientKey + "|" + TimeStamp
// @Param	body	body	dtos.RegisterPi3Request	true	"Device Information"
// @Success 200 {object} dtos.RegisterPi3Response
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /register-pi3 [post]
func (c *Pi3UpdateInfoController) RegisterPi3(body *dtos.RegisterPi3Request) {
	if body.PosID == "" {
		logger.Errorf(c.Ctx.Request.Context(), "PosID invalid")
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
	dataInput := fmt.Sprintf("%s|%s|%d", body.PosID, cfgutil.Load("TNG_CLIENT_KEY"), body.TimeStamp)
	sig := hashutil.GetSHA256(dataInput)
	if sig != body.Sig && runMode != beego.DEV {
		logger.Errorf(c.Ctx.Request.Context(), "Sig not matching")
		c.Respond(nil, dtos.NewAppError(dtos.ErrorSigNotMatching))
		return
	}

	c.Respond(c.pi3UpdateInfoService.RegisterPi3(c.Ctx.Request.Context(), body))
}
