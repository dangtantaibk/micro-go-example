package controllers

import (
	"tng/common/logger"
	jwtUtils "tng/common/utils/jwt"
	"tng/common/utils/strutil"
	"tng/shipper-service/dtos"
	"tng/shipper-service/services"
)


// InvoiceController represents controller of invoice
type InvoiceController struct {
	BaseController
	invoiceService services.InvoiceService
}

// Prepare handles prepare of InvoiceController.
func (c *InvoiceController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.InvoiceService) {
		c.invoiceService = s
	})
}

// Get List Invoice.
// @Title Get List Invoice.
// @Description Get List Invoice
// @Param date	query	string	true	"Date format yyyy-MM-dd"
// @Param merchant_code	query	string	true	"Merchant code"
// @Param current_page	query	int32	false	"Integer"
// @Param total_trans_per_page	query	int32	false	"Integer"
// @Success 200 {object} dtos.ListInvoiceResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /list [get]
func (c *InvoiceController) List() {
	var request dtos.ListInvoiceRequest

	authHeader := c.Ctx.Input.Header("Authorization")
	tokenString := jwtUtils.ParseAuthHeaderToken(authHeader)

	if strutil.IsEmpty(tokenString) {
		logger.Errorf(c.Ctx.Request.Context(), "Invalid token.")
		c.Respond(nil, dtos.NewAppError(dtos.UnauthorizedError))
		return
	}

	claims, err := jwtUtils.ParseWithClaims(tokenString)

	if err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Couldn't handle this token:", err)
		c.Respond(nil, dtos.NewAppError(dtos.UnauthorizedError))
		return
	}

	phoneNumber := claims.(*dtos.Claims).PhoneNumber
	if strutil.IsEmpty(phoneNumber) {
		logger.Errorf(c.Ctx.Request.Context(), "Phone number is empty")
		c.Respond(nil, dtos.NewAppError(dtos.UnauthorizedError))
		return
	}

	if err := c.ParseForm(&request); err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Parsing form: %v", err)
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	if request.Date == "" {
		logger.Errorf(c.Ctx.Request.Context(), "Date invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.MerchantCode == "" {
		logger.Errorf(c.Ctx.Request.Context(), "MerchantCode invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	c.Respond(c.invoiceService.List(c.Ctx.Request.Context(), &request, phoneNumber))
}

// Get All List Invoice.
// @Title Get List Invoice.
// @Description Get List Invoice
// @Param date	query	string	true	"Date format yyyy-MM-dd"
// @Param merchant_code	query	string	true	"Merchant code"
// @Param current_page	query	int32	false	"Integer"
// @Param total_trans_per_page	query	int32	false	"Integer"
// @Success 200 {object} dtos.ListInvoiceResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /all-list [get]
func (c *InvoiceController) AllList() {
	var request dtos.ListInvoiceRequest

	authHeader := c.Ctx.Input.Header("Authorization")
	tokenString := jwtUtils.ParseAuthHeaderToken(authHeader)

	if strutil.IsEmpty(tokenString) {
		logger.Errorf(c.Ctx.Request.Context(), "Invalid token.")
		c.Respond(nil, dtos.NewAppError(dtos.UnauthorizedError))
		return
	}

	claims, err := jwtUtils.ParseWithClaims(tokenString)

	if err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Couldn't handle this token:", err)
		c.Respond(nil, dtos.NewAppError(dtos.UnauthorizedError))
		return
	}

	phoneNumber := claims.(*dtos.Claims).PhoneNumber
	if strutil.IsEmpty(phoneNumber) {
		logger.Errorf(c.Ctx.Request.Context(), "Phone number is empty")
		c.Respond(nil, dtos.NewAppError(dtos.UnauthorizedError))
		return
	}

	if err := c.ParseForm(&request); err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Parsing form: %v", err)
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	if request.Date == "" {
		logger.Errorf(c.Ctx.Request.Context(), "Date invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.MerchantCode == "" {
		logger.Errorf(c.Ctx.Request.Context(), "MerchantCode invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	c.Respond(c.invoiceService.AllList(c.Ctx.Request.Context(), &request))
}

// Update payment status
// @Title Update Payment Status.
// @Description Update Payment Status.
// @Param	body	body	true	dtos.UpdateInvoiceStatusRequest
// @Success 200 {object} dtos.UpdateInvoiceStatusResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /update-status [post]
func (c *InvoiceController) UpdateStatus(body *dtos.UpdateInvoiceStatusRequest) {
	result, err := c.isTokenVerified()

	if err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Couldn't handle this token:", err)
		c.Respond(nil, dtos.NewAppError(dtos.UnauthorizedError))
		return
	}

	if result == 2 {
		logger.Errorf(c.Ctx.Request.Context(), "Invalid token.")
		c.Respond(nil, dtos.NewAppError(dtos.UnauthorizedError))
		return
	}

	if body.InvoiceCode == "" {
		logger.Errorf(c.Ctx.Request.Context(), "Invoice code invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if body.MerchantCode == "" {
		logger.Errorf(c.Ctx.Request.Context(), "MerchantCode invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	c.Respond(c.invoiceService.UpdateStatus(c.Ctx.Request.Context(), body))

}

// Get Invoice and Invoice Detail.
// @Title Get Invoice and Invoice Detail.
// @Description Get Invoice and Invoice Detail by vpos token
// @Param merchant_code	query	string	true	"Merchant code"
// @Param vpos_token	query	string	true	"Vpos token"
// @Success 200 {object} dtos.ScanQRCodeResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /scan-qr-code [get]
func (c *InvoiceController) ScanQRCode() {
	result, err := c.isTokenVerified()

	if err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Couldn't handle this token:", err)
		c.Respond(nil, dtos.NewAppError(dtos.UnauthorizedError))
		return
	}

	if result == 2 {
		logger.Errorf(c.Ctx.Request.Context(), "Invalid token.")
		c.Respond(nil, dtos.NewAppError(dtos.UnauthorizedError))
		return
	}

	var request dtos.ScanQRCodeRequest
	if err := c.ParseForm(&request); err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Parsing form: %v", err)
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.MerchantCode == "" {
		logger.Errorf(c.Ctx.Request.Context(), "MerchantCode invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.VposToken == "" {
		logger.Errorf(c.Ctx.Request.Context(), "VposToken invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	c.Respond(c.invoiceService.GetByVposToken(c.Ctx.Request.Context(), &request))
}

// Get Invoice Detail.
// @Title Get Invoice Detail.
// @Description Get Invoice Detail by invoice code
// @Param merchant_code	query	string	true	"Merchant code"
// @Param invoice_code	query	string	true	"Invoice code"
// @Success 200 {object} dtos.GetInvoiceDetailResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /invoice-detail [get]
func (c *InvoiceController) GetInvoiceDetail() {
	result, err := c.isTokenVerified()

	if err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Couldn't handle this token:", err)
		c.Respond(nil, dtos.NewAppError(dtos.UnauthorizedError))
		return
	}

	if result == 2 {
		logger.Errorf(c.Ctx.Request.Context(), "Invalid token.")
		c.Respond(nil, dtos.NewAppError(dtos.UnauthorizedError))
		return
	}

	var request dtos.GetInvoiceDetailRequest
	if err := c.ParseForm(&request); err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Parsing form: %v", err)
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.MerchantCode == "" {
		logger.Errorf(c.Ctx.Request.Context(), "MerchantCode invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}
	if request.InvoiceCode == "" {
		logger.Errorf(c.Ctx.Request.Context(), "Invoice code invalid")
		c.Respond(nil, dtos.NewAppError(dtos.InvalidRequestError))
		return
	}

	c.Respond(c.invoiceService.GetInvoiceDetail(c.Ctx.Request.Context(), &request))
}


func (c *InvoiceController) isTokenVerified() (int, error) {
	authHeader := c.Ctx.Input.Header("Authorization")
	return jwtUtils.TokenVerify(authHeader)
}

