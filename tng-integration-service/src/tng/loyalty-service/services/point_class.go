
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

type PointClassService interface {
	InsertOrUpdate(context.Context, *dtos.InsertOrUpdatePointClassRequest) (*dtos.InsertOrUpdatePointClassResponse, error)
	Delete(context.Context, *dtos.DeletePointClassRequest) (*dtos.DeletePointClassResponse, error)
	GetByID(context.Context, *dtos.GetPointClassRequest) (*dtos.GetPointClassResponse, error)
	List(context.Context, *dtos.ListPointClassRequest) (*dtos.ListPointClassResponse, error)
}

type pointClassService struct {
	BaseService
	redisCache        redisutil.Cache
	pointClassRepository repositories.PointClassRepository
}

func NewPointClassService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	pointClassRepository repositories.PointClassRepository,
) PointClassService {
	return &pointClassService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:        redisCache,
		pointClassRepository: pointClassRepository,
	}
}

func (s *pointClassService) InsertOrUpdate(ctx context.Context, request *dtos.InsertOrUpdatePointClassRequest) (*dtos.InsertOrUpdatePointClassResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	modelPointClass := &loyalty.PointClass{

            
            Id: request.Id,

                        
            Title: request.Title,

            
            MinAccMoney: request.MinAccMoney,

            
            DiscountPercent: request.DiscountPercent,

            
            RequireNumOfTrans: request.RequireNumOfTrans,

            
            RequireNumOfPoint: request.RequireNumOfPoint,

            
            JsonDetail: request.JsonDetail,

            
            Created: request.Created,

            	}
	err := s.pointClassRepository.InsertOrUpdate(ctx, tx, modelPointClass)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "InsertOrUpdate pointClass error: %v", err)
	}
	resp := &dtos.InsertOrUpdatePointClassResponse{
		Meta: dtos.NewMetaOK(),
	}
	return resp, nil
}

func (s *pointClassService) Delete(ctx context.Context, request *dtos.DeletePointClassRequest) (*dtos.DeletePointClassResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	err := s.pointClassRepository.Delete(ctx, tx, request.Id)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Delete pointClass error: %v", err)
		return nil, err
	}
	resp := &dtos.DeletePointClassResponse{
		Meta: dtos.NewMetaOK(),
	}
	return resp, nil
}

func (s *pointClassService) GetByID(ctx context.Context, request *dtos.GetPointClassRequest) (*dtos.GetPointClassResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	item, err := s.pointClassRepository.GetByID(ctx, tx, request.Id)
	if err != nil {
		logger.Errorf(ctx, "GetByID pointClass error: %v", err)
		return nil, err
	}
	data := &dtos.PointClassInfo{}
	err = copier.Copy(&data, &item)
	if err != nil {
		logger.Errorf(ctx, "Copy pointClass error: %v", err)
		return nil, err
	}
	resp := &dtos.GetPointClassResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}

func (s *pointClassService) List(ctx context.Context, request *dtos.ListPointClassRequest) (*dtos.ListPointClassResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	list, err := s.pointClassRepository.List(ctx, tx, request.PageIndex, request.PageSize)
	if err != nil {
		logger.Errorf(ctx, "List pointClass error: %v", err)
		return nil, err
	}
	data := make([]*dtos.PointClassInfo, 0)
	for _, item := range list {
		var (
			iv dtos.PointClassInfo
			_  = copier.Copy(&iv, &item)
		)
		data = append(data, &iv)
	}
	resp := &dtos.ListPointClassResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}

