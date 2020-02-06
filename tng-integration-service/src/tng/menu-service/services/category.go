package services

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"time"
	"tng/common/logger"
	"tng/common/models"
	"tng/common/models/menu"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/menu-service/dtos"
	"tng/menu-service/repositories"
)

type categoryService struct {
	BaseService
	redisCache redisutil.Cache
	categoryRepository repositories.CategoryRepository
}

func (c *categoryService) Delete(ctx context.Context, request *dtos.DeleteCategoryRequest) (*dtos.DeleteCategoryResponse, error) {
	var (
		tx = c.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)
	defer c.dbFactory.Rollback(tx)
	err := c.categoryRepository.Delete(tx, request.CategoryID)
	c.dbFactory.Commit(tx)

	if err != nil {
		logger.Errorf(ctx, "Delele category error: %v", err)
		return nil, dtos.NewAppError(dtos.InternalServerError)
	}

	response := &dtos.DeleteCategoryResponse{Meta: dtos.Meta{
		Code:    1,
		Message: "OK",
	}}

	return response, nil
}

func (c *categoryService) Update(ctx context.Context, request *dtos.UpdateCategoryRequest) (*dtos.UpdateCategoryResponse, error) {
	var (
		tx = c.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
		categoryModel = &menu.Category{
			ID:               request.CategoryID,
			CategoryName:     request.CategoryName,
			CategoryValue:    0,
			CategoryType:     0,
			Order:            request.Order,
			Status:           request.Status,
			AreaID:           request.AreaID,
			ModifiedBy:       "",
			ModifiedDateTime: time.Now().Local(),
		}
	)

	defer c.dbFactory.Rollback(tx)
	err := c.categoryRepository.Update(tx, categoryModel)
	c.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Update category error: %v", err)
		return nil, err
	}
	response := &dtos.UpdateCategoryResponse{Meta:dtos.Meta{
		Code:    1,
		Message: "OK",
	}}

	return response, nil
}

func (c *categoryService) UpdateStatus(ctx context.Context, request *dtos.UpdateStatusRequest) (*dtos.UpdateStatusResponse, error) {
	var (
		tx = c.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)
	defer c.dbFactory.Rollback(tx)
	err := c.categoryRepository.UpdateStatus(tx, request.CategoryID, request.Status)
	c.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Update status category error: %v", err)
		return nil, dtos.NewAppError(dtos.InternalServerError)
	}

	response := &dtos.UpdateStatusResponse{Meta:dtos.Meta{
		Code:    1,
		Message: "OK",
	}}

	return response, nil
}

func (c *categoryService) List(ctx context.Context, request *dtos.ListCategoryRequest) (*dtos.ListCategoryResponse, error) {
	var (
		tx = c.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)

	list, err := c.categoryRepository.List(tx)
	if err != nil {
		logger.Errorf(ctx, "Get list category error: %v", err)
		return nil, err
	}

	data := make([]dtos.Category, 0)
	for _, item := range list {
		var (
			cat dtos.Category
			_ = copier.Copy(&cat, &item)
		)
		data = append(data, cat)
	}

	response := &dtos.ListCategoryResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
		Data: data,
	}
	return response, nil
}

func (c *categoryService) Insert(ctx context.Context, request *dtos.CreateCategoryRequest) (*dtos.CreateCategoryResponse, error) {
	var (
		tx = c.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
		modelCategory = &menu.Category{
			CategoryName:     request.CategoryName,
			CategoryValue:    0,
			CategoryType:     0,
			Order:            request.Order,
			Status:           request.Status,
			AreaID:           request.AreaID,
			ModifiedBy:       "",
			ModifiedDateTime: time.Now(),
		}
	)

	defer c.dbFactory.Rollback(tx)
	err := c.categoryRepository.Insert(tx, modelCategory)
	c.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Insert category error: %v", err)
		return nil, err
	}

	response := &dtos.CreateCategoryResponse{Meta:dtos.Meta{
		Code:    1,
		Message: "OK",
	}}

	return response, nil
}

type CategoryService interface {
	List(ctx context.Context, request *dtos.ListCategoryRequest) (*dtos.ListCategoryResponse, error)
	Insert(ctx context.Context, request *dtos.CreateCategoryRequest) (*dtos.CreateCategoryResponse, error)
	UpdateStatus(ctx context.Context, request *dtos.UpdateStatusRequest) (*dtos.UpdateStatusResponse, error)
	Delete(ctx context.Context, request *dtos.DeleteCategoryRequest) (*dtos.DeleteCategoryResponse, error)
	Update(ctx context.Context, request *dtos.UpdateCategoryRequest) (*dtos.UpdateCategoryResponse, error)
}

func NewCategoryService(factory db.Factory, redisCache redisutil.Cache, categoryRepository repositories.CategoryRepository) CategoryService {
	return &categoryService{
		BaseService:        BaseService{dbFactory:factory},
		redisCache:         redisCache,
		categoryRepository: categoryRepository,
	}
}