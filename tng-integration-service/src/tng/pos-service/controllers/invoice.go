package controllers

import (
	"tng/pos-service/dtos"
	"tng/pos-service/services"
)

var (
//runMode = cfgutil.Load("RUNMODE")
)

// InvoiceController represents controller of Token
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

// Create Invoice.
// @Title Create Invoice
// @Description Call to ZaloPay Create Invoice
// @Param	body	body	dtos.CreateInvoiceRequest	true	"Information Order"
// @Success 200 {object} dtos.CreateInvoiceResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /create [post]
func (c *InvoiceController) Create(body *dtos.CreateInvoiceRequest) {
	c.Respond(c.invoiceService.Create(c.Ctx.Request.Context(), body))
}

// Cancel Invoice.
// @Title Cancel Invoice
// @Description Cancel Invoice
// @Param	body	body	dtos.CancelInvoiceRequest	true	"Information Order"
// @Success 200 {object} dtos.CancelInvoiceResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /cancel [post]
func (c *InvoiceController) Cancel(body *dtos.CancelInvoiceRequest) {
	c.Respond(c.invoiceService.Cancel(c.Ctx.Request.Context(), body))
}

// Refund Invoice.
// @Title Refund Invoice
// @Description Refund Invoice
// @Param	body	body	dtos.RefundInvoiceRequest	true	"Information Refund"
// @Success 200 {object} dtos.RefundInvoiceResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /refund [post]
func (c *InvoiceController) Refund(body *dtos.RefundInvoiceRequest) {
	c.Respond(c.invoiceService.Refund(c.Ctx.Request.Context(), body))
}
