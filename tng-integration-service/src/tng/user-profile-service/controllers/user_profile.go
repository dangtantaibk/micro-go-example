package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"tng/common/logger"
	"tng/common/utils/hashutil"
	"tng/common/utils/keysutil"
	"tng/user-profile-service/dtos"
	"tng/user-profile-service/services"
)

// UserProfileController represents controller
type UserProfileController struct {
	BaseController
	userService services.UserProfileService
}

// Prepare handles prepare of UserController.
func (c *UserProfileController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.UserProfileService) {
		c.userService = s
	})
}

// Create.
// @Title Create
// @Description Create
// @Param	body	body		dtos.CreateProfileRequest	true	"Information Request"
// @Success 200 	{object} 	dtos.CreateProfileResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /create [post]
func (c *UserProfileController) Create(body *dtos.CreateProfileRequest) {
	if body == nil {
		logger.Errorf(c.Ctx.Request.Context(), "invalid request")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	c.Respond(c.userService.Create(c.Ctx.Request.Context(), body))
}

// Update.
// @Title Update
// @Description Update
// @Param	body	body		dtos.UpdateProfileRequest	true	"Information Request"
// @Success 200 	{object} 	dtos.UpdateProfileResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /update [post]
func (c *UserProfileController) Update(body *dtos.UpdateProfileRequest) {
	if body == nil {
		logger.Errorf(c.Ctx.Request.Context(), "invalid request")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	c.Respond(c.userService.Update(c.Ctx.Request.Context(), body))
}

// GetByID.
// @Title GetByID
// @Description Get
// @Param user_id	query	string	true	"UserID for getting."
// @Param app_id	query	int		true	"AppID for getting."
// @Param timestamp	query	int		true	"TimeStamp for current request. Ex: 1572849166"
// @Param sig		query	string	true	"Sig for SHA256. input: AppId + "|" + ClientKey + "|" + TimeStamp"
// @Success 200 {object} dtos.GetProfileResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /get-by-id [get]
func (c *UserProfileController) GetByID() {
	var request dtos.GetProfileByIDRequest
	if err := c.ParseForm(&request); err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Parsing form: %v", err)
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	if request.AppID <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "AppID invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.UserID <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "UserID invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.Timestamp <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "Timestamp invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.Sig == "" {
		logger.Errorf(c.Ctx.Request.Context(), "Sig invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	clientKey := keysutil.GetHashKey(fmt.Sprintf("%d", request.AppID))
	dataInput := fmt.Sprintf("%d|%s|%d", request.AppID, clientKey, request.Timestamp)
	sig := hashutil.GetSHA256(dataInput)
	if sig != request.Sig && runMode != beego.DEV {
		logger.Errorf(c.Ctx.Request.Context(), "Sig not matching")
		c.Respond(nil, dtos.NewAppError(dtos.MatchingSigError))
		return
	}

	c.Respond(c.userService.GetByID(c.Ctx.Request.Context(), &request))
}
