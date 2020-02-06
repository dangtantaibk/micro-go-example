package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"tng/common/logger"
	"tng/common/utils/hashutil"
	"tng/common/utils/keysutil"
	"tng/promotion-service/dtos"
	"tng/promotion-service/services"
)

// CampaignController represents controller of Token
type CampaignController struct {
	BaseController
	campaignService services.CampaignService
}

// Prepare handles prepare of H5ZaloPayController.
func (c *CampaignController) Prepare() {
	//c.BaseController.Prepare()
	_ = services.GetServiceContainer().Invoke(func(s services.CampaignService) {
		c.campaignService = s
	})
}

// Insert campaign.
// @Title Insert campaign
// @Description Insert campaign, Sig for SHA256. input: AppID + "|" + clientKey + "|" + TimeStamp
// @Param	body	body	dtos.InsertCampaignRequest	true	"Information request"
// @Success 200 {object} dtos.InsertCampaignResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /insert [post]
func (c *CampaignController) Insert(body *dtos.InsertCampaignRequest) {
	clientKey := keysutil.GetKey(fmt.Sprintf("%d", body.AppID))
	dataInput := fmt.Sprintf("%d|%s|%d", body.AppID, clientKey, body.TimeStamp)
	sig := hashutil.GetSHA256(dataInput)
	if sig != body.Sig && runMode != beego.DEV {
		logger.Errorf(c.Ctx.Request.Context(), "Sig not matching")
		c.Respond(nil, dtos.NewAppError(dtos.ErrorSigNotMatching))
		return
	}

	c.Respond(c.campaignService.Insert(c.Ctx.Request.Context(), body))
}


// Get List Campaign.
// @Title Get List Campaign.
// @Description Get List Campaign
// @Param appid			query	string	true	"App ID"
// @Param page_index	query	int32	true	"Page Index"
// @Param page_size	  	query	int32	true	"Page Size"
// @Success 200 {object} dtos.ListCampaignResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /list [get]
func (c *CampaignController) List() {
	var request dtos.ListCampaignRequest
	if err := c.ParseForm(&request); err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Parsing form: %v", err)
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	if request.AppID == "" {
		logger.Errorf(c.Ctx.Request.Context(), "AppID invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.PageSize <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "PageSize invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.PageIndex <= 0 {
		logger.Errorf(c.Ctx.Request.Context(), "PageIndex invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	c.Respond(c.campaignService.List(c.Ctx.Request.Context(), &request))
}
