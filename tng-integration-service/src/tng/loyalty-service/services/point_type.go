
package services

import (
	"context"
	"github.com/jinzhu/copier"
	"tng/common/logger"
	"tng/common/models/loyalty"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/loyalty-service/dtos"
	"tng/loyalty-service/repositories"
)

type PointTypeService interface {
	InsertOrUpdate(context.Context, *dtos.InsertOrUpdatePointTypeRequest) (*dtos.InsertOrUpdatePointTypeResponse, error)
	Delete(context.Context, *dtos.DeletePointTypeRequest) (*dtos.DeletePointTypeResponse, error)
	GetByID(context.Context, *dtos.GetPointTypeRequest) (*dtos.GetPointTypeResponse, error)
	List(context.Context, *dtos.ListPointTypeRequest) (*dtos.ListPointTypeResponse, error)
}

type pointTypeService struct {
	BaseService
	redisCache        redisutil.Cache
	pointTypeRepository repositories.PointTypeRepository
}

func NewPointTypeService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	pointTypeRepository repositories.PointTypeRepository,
) PointTypeService {
	return &pointTypeService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:        redisCache,
		pointTypeRepository: pointTypeRepository,
	}
}

func (s *pointTypeService) InsertOrUpdate(ctx context.Context, request *dtos.InsertOrUpdatePointTypeRequest) (*dtos.InsertOrUpdatePointTypeResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	modelPointType := &loyalty.PointType{

            
            Id: request.Id,

                        
            Description: request.Description,

            
            JsonDetail: request.JsonDetail,

            
            Created: request.Created,

            	}
	err := s.pointTypeRepository.InsertOrUpdate(ctx, tx, modelPointType)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "InsertOrUpdate pointType error: %v", err)
	}
	resp := &dtos.InsertOrUpdatePointTypeResponse{
		Meta: dtos.NewMetaOK(),
	}
	return resp, nil
}

func (s *pointTypeService) Delete(ctx context.Context, request *dtos.DeletePointTypeRequest) (*dtos.DeletePointTypeResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	err := s.pointTypeRepository.Delete(ctx, tx, request.Id)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Delete pointType error: %v", err)
		return nil, err
	}
	resp := &dtos.DeletePointTypeResponse{
		Meta: dtos.NewMetaOK(),
	}
	return resp, nil
}

func (s *pointTypeService) GetByID(ctx context.Context, request *dtos.GetPointTypeRequest) (*dtos.GetPointTypeResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	item, err := s.pointTypeRepository.GetByID(ctx, tx, request.Id)
	if err != nil {
		logger.Errorf(ctx, "GetByID pointType error: %v", err)
		return nil, err
	}
	data := &dtos.PointTypeInfo{}
	err = copier.Copy(&data, &item)
	if err != nil {
		logger.Errorf(ctx, "Copy pointType error: %v", err)
		return nil, err
	}
	resp := &dtos.GetPointTypeResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}

func (s *pointTypeService) List(ctx context.Context, request *dtos.ListPointTypeRequest) (*dtos.ListPointTypeResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	list, err := s.pointTypeRepository.List(ctx, tx, request.PageIndex, request.PageSize)
	if err != nil {
		logger.Errorf(ctx, "List pointType error: %v", err)
		return nil, err
	}
	data := make([]*dtos.PointTypeInfo, 0)
	for _, item := range list {
		var (
			iv dtos.PointTypeInfo
			_  = copier.Copy(&iv, &item)
		)
		data = append(data, &iv)
	}
	resp := &dtos.ListPointTypeResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}

