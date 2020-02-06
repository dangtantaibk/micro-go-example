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

// PromotionController represents controller of Token
type PromotionController struct {
	BaseController
	promotionService services.PromotionService
}

// Prepare handles prepare of H5ZaloPayController.
func (c *PromotionController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.PromotionService) {
		c.promotionService = s
	})
}

// Insert Promotion.
// @Title Insert Promotion
// @Description Insert Promotion, Sig for SHA256. input: AppID + "|" + clientKey + "|" + TimeStamp
// @Param	body	body	dtos.PromotionRequest	true	"Information request"
// @Success 200 {object} dtos.PromotionResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /insert [post]
func (c *PromotionController) Insert(body *dtos.PromotionRequest) {
	clientKey := keysutil.GetKey(body.AppID)
	dataInput := fmt.Sprintf("%s|%s|%d", body.AppID, clientKey, body.TS)
	sig := hashutil.GetSHA256(dataInput)
	if sig != body.CheckSum && runMode != beego.DEV {
		logger.Errorf(c.Ctx.Request.Context(), "Sig not matching")
		c.Respond(nil, dtos.NewAppError(dtos.ErrorSigNotMatching))
		return
	}

	c.Respond(c.promotionService.Insert(c.Ctx.Request.Context(), body))
}
