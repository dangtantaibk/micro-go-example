package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"tng/common/logger"
	"tng/common/utils/hashutil"
	"tng/common/utils/keysutil"
	"tng/user-profile-service/dtos"
	"tng/user-profile-service/services"
)

// AuthenticationController represents controller
type AuthenticationController struct {
	BaseController
	userSessionService services.AuthenticationService
}

// Prepare handles prepare of UserSessionController.
func (c *AuthenticationController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.AuthenticationService) {
		c.userSessionService = s
	})
}

// Login.
// @Title Login
// @Description Login
// @Param	body	body		dtos.LoginRequest	true	"Information Request For Login"
// @Success 200 	{object} 	dtos.LoginResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /login [post]
func (c *AuthenticationController) Login(body *dtos.LoginRequest) {
	// check data request
	if body == nil {
		logger.Errorf(c.Ctx.Request.Context(), "invalid request")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if body.LoginType <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "login type invalid")
		c.Respond(nil, dtos.NewAppError(dtos.LoginInfoInvalid))
		return
	}

	// check data login
	dataRequest := body.Data
	if dataRequest == "" {
		logger.Errorf(c.Ctx.Request.Context(), "data invalid")
		c.Respond(nil, dtos.NewAppError(dtos.LoginInfoInvalid))
		return
	}
	dataLogin := &dtos.DataLoginRequest{}
	err := json.Unmarshal([]byte(dataRequest), dataLogin)
	if err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "unmarshal data login error: %v", err)
		c.Respond(nil, dtos.NewAppError(dtos.LoginInfoInvalid))
		return
	}
	if dataLogin.SocialID == "" || dataLogin.Platform == "" || dataLogin.AppID <= 0 || dataLogin.TS <= 0 || dataLogin.Sig == "" {
		logger.Errorf(c.Ctx.Request.Context(), "invalid invalid")
		c.Respond(nil, dtos.NewAppError(dtos.LoginInfoInvalid))
		return
	}

	// check sig
	clientKey := keysutil.GetHashKey(fmt.Sprintf("%d", dataLogin.AppID))
	dataInput := fmt.Sprintf("%s|%d|%s|%d|%s", dataLogin.SocialID, dataLogin.AppID, dataLogin.Platform, dataLogin.TS, clientKey)
	sig := hashutil.GetSHA256(dataInput)
	if sig != dataLogin.Sig && runMode != beego.DEV {
		logger.Errorf(c.Ctx.Request.Context(), "Sig not matching")
		c.Respond(nil, dtos.NewAppError(dtos.MatchingSigError))
		return
	}

	c.Respond(c.userSessionService.Login(c.Ctx.Request.Context(), dataLogin))
}

// CheckLogin.
// @Title CheckLogin
// @Description CheckLogin
// @Param	body	body		dtos.CheckLoginRequest	true	"Information Request For Check Login"
// @Success 200 	{object} 	dtos.CheckLoginResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /check-login [post]
func (c *AuthenticationController) CheckLogin(body *dtos.CheckLoginRequest) {
	if body == nil {
		logger.Errorf(c.Ctx.Request.Context(), "invalid request")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if body.Token == "" {
		logger.Errorf(c.Ctx.Request.Context(), "token invalid")
		c.Respond(nil, dtos.NewAppError(dtos.LoginInfoInvalid))
		return
	}
	c.Respond(c.userSessionService.CheckLogin(c.Ctx.Request.Context(), body))
}
