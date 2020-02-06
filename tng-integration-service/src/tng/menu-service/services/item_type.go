package services

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"tng/common/logger"
	"tng/common/models"
	"tng/common/models/menu"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/menu-service/dtos"
	"tng/menu-service/repositories"
)

type itemTypeService struct {
	BaseService
	redisCache         redisutil.Cache
	itemTypeRepository repositories.ItemTypeRepository
}

func (a *itemTypeService) Create(ctx context.Context, request *dtos.CreateItemTypeRequest) (*dtos.CreateItemTypeResponse, error) {
	var (
		tx            = a.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
		itemTypeModel = &menu.ItemType{
			ItemTypeName: request.ItemTypeName,
			Status:       request.Status,
		}
	)
	defer a.dbFactory.Rollback(tx)
	itemTypeInfo, err := a.itemTypeRepository.Insert(tx, itemTypeModel)
	a.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Create item type error: %v", err)
		return nil, err
	}
	var itemType dtos.ItemType
	err = copier.Copy(&itemType, &itemTypeInfo)
	if err != nil {
		logger.Errorf(ctx, "parse item type error: %v", err)
		return nil, err
	}

	response := &dtos.CreateItemTypeResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
		Data: itemType,
	}
	return response, nil
}

func (a *itemTypeService) Delete(ctx context.Context, request *dtos.DeleteItemTypeRequest) (*dtos.DeleteItemTypeResponse, error) {
	var (
		tx = a.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)
	defer a.dbFactory.Rollback(tx)
	err := a.itemTypeRepository.Delete(tx, request.ItemTypeID)
	a.dbFactory.Commit(tx)

	if err != nil {
		logger.Errorf(ctx, "Delele item type error: %v", err)
		return nil, dtos.NewAppError(dtos.InternalServerError)
	}
	response := &dtos.DeleteItemTypeResponse{Meta: dtos.Meta{
		Code:    1,
		Message: "OK",
	},}
	return response, nil
}

func (a *itemTypeService) Update(ctx context.Context, request *dtos.UpdateItemTypeRequest) (*dtos.UpdateItemTypeResponse, error) {
	var (
		tx            = a.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
		itemTypeModel = &menu.ItemType{
			ID:           request.ItemTypeID,
			ItemTypeName: request.ItemTypeName,
			Status:       request.Status,
		}
	)
	defer a.dbFactory.Rollback(tx)
	itemTypeInfo, err := a.itemTypeRepository.Update(tx, itemTypeModel)
	a.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Update item type error: %v", err)
		return nil, dtos.NewAppError(dtos.InternalServerError)
	}
	var itemType dtos.ItemType
	err = copier.Copy(&itemType, &itemTypeInfo)
	if err != nil {
		logger.Errorf(ctx, "parse item type error: %v", err)
		return nil, err
	}

	response := &dtos.UpdateItemTypeResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
		Data: itemType,
	}
	return response, nil
}

func (a *itemTypeService) List(ctx context.Context, request *dtos.ListItemTypeRequest) (*dtos.ListItemTypeResponse, error) {
	var (
		tx = a.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)

	list, err := a.itemTypeRepository.List(tx)
	if err != nil {
		logger.Errorf(ctx, "Get list area error: %v", err)
		return nil, err
	}

	data := make([]dtos.ItemType, 0)
	for _, item := range list {
		var (
			itemType dtos.ItemType
			_        = copier.Copy(&itemType, &item)
		)
		data = append(data, itemType)
	}

	response := &dtos.ListItemTypeResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
		Data: data,
	}
	return response, nil
}

type ItemTypeService interface {
	List(ctx context.Context, request *dtos.ListItemTypeRequest) (*dtos.ListItemTypeResponse, error)
	Create(ctx context.Context, request *dtos.CreateItemTypeRequest) (*dtos.CreateItemTypeResponse, error)
	Delete(ctx context.Context, request *dtos.DeleteItemTypeRequest) (*dtos.DeleteItemTypeResponse, error)
	Update(ctx context.Context, request *dtos.UpdateItemTypeRequest) (*dtos.UpdateItemTypeResponse, error)
}

func NewItemTypeService(factory db.Factory, cache redisutil.Cache, itemTypeRepository repositories.ItemTypeRepository) ItemTypeService {
	return &itemTypeService{
		BaseService:        BaseService{dbFactory: factory},
		redisCache:         cache,
		itemTypeRepository: itemTypeRepository,
	}
}
