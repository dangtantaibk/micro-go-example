package controllers

import (
	"tng/common/logger"
	"tng/menu-service/dtos"
	"tng/menu-service/services"
)

// CategoryController represent controller of category
type CategoryController struct {
	BaseController
	categoryService services.CategoryService
}

// Prepare handle of CategoryController
func (c *CategoryController) Prepare() {
	//c.BaseController.Prepare()
	_ = services.GetServiceContainer().Invoke(func(s services.CategoryService) {
		c.categoryService = s
	})
}

// Get List Category.
// @Title Get List Category.
// @Description Get List Category
// @Param merchant_code	query string true "Merchant code"
// @Success 200 {object} dtos.ListCategoryResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /list [get]
func (c *CategoryController) List() {
	var request dtos.ListCategoryRequest
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
	c.Respond(c.categoryService.List(c.Ctx.Request.Context(), &request))
}

// Create Category.
// @Title Create Category
// @Description Create Category
// @Param	body	body	dtos.CreateCategoryRequest	true	"Category info"
// @Success 200 {object} dtos.CreateCategoryResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /create [post]
func (c *CategoryController) Create(body *dtos.CreateCategoryRequest) {
	c.Respond(c.categoryService.Insert(c.Ctx.Request.Context(), body))
}

// Update Category status
// @Title Update Category Status.
// @Description Update Category Status.
// @Param	body	body dtos.UpdateStatusRequest	true	"Update request info"
// @Success 200 {object} dtos.UpdateStatusResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /update-status [post]
func (c *CategoryController) UpdateStatus(body *dtos.UpdateStatusRequest) {
	c.Respond(c.categoryService.UpdateStatus(c.Ctx.Request.Context(), body))
}

// Delete Category
// @Title Delete Category
// @Description Delete Category
// @Param	body	body dtos.DeleteCategoryRequest	true	"Update request info"
// @Success 200 {object} dtos.DeleteCategoryResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /delete [post]
func (c *CategoryController) DeleteCategory(body *dtos.DeleteCategoryRequest) {
	c.Respond(c.categoryService.Delete(c.Ctx.Request.Context(), body))
}

// Update Category
// @Title Update Category
// @Description Update Category
// @Param	body	body dtos.UpdateCategoryRequest	true	"Update request info"
// @Success 200 {object} dtos.UpdateCategoryResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /update [post]
func (c *CategoryController) Update(body *dtos.UpdateCategoryRequest) {
	c.Respond(c.categoryService.Update(c.Ctx.Request.Context(), body))
}
