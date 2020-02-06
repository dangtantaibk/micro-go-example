package services

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"tng/common/logger"
	"tng/common/models"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/menu-service/dtos"
	"tng/menu-service/repositories"
)

type areaService struct {
	BaseService
	redisCache redisutil.Cache
	areaRepository repositories.AreaRepository
}

func (a *areaService) List(ctx context.Context, request *dtos.ListAreaRequest) (*dtos.ListAreaResponse, error) {
	var (
		tx = a.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)

	list, err := a.areaRepository.List(tx)
	if err != nil {
		logger.Errorf(ctx, "Get list area error: %v", err)
		return nil, err
	}

	data := make([]dtos.Area , 0)
	for _, item := range list {
		var (
			area dtos.Area
			_ = copier.Copy(&area, &item)
		)
		data = append(data, area)
	}

	response :=  &dtos.ListAreaResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
		Data: data,
	}
	return response, nil
}

type AreaService interface {
	List(ctx context.Context, request *dtos.ListAreaRequest) (*dtos.ListAreaResponse, error)
}

func NewAreaService(factory db.Factory, cache redisutil.Cache, areaRepository repositories.AreaRepository) AreaService {
	return &areaService{
		BaseService:    BaseService{dbFactory:factory},
		redisCache:     cache,
		areaRepository: areaRepository,
	}
}
