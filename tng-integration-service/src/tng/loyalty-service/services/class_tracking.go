
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

type ClassTrackingService interface {
	InsertOrUpdate(context.Context, *dtos.InsertOrUpdateClassTrackingRequest) (*dtos.InsertOrUpdateClassTrackingResponse, error)
	Delete(context.Context, *dtos.DeleteClassTrackingRequest) (*dtos.DeleteClassTrackingResponse, error)
	GetByID(context.Context, *dtos.GetClassTrackingRequest) (*dtos.GetClassTrackingResponse, error)
	List(context.Context, *dtos.ListClassTrackingRequest) (*dtos.ListClassTrackingResponse, error)
}

type classTrackingService struct {
	BaseService
	redisCache        redisutil.Cache
	classTrackingRepository repositories.ClassTrackingRepository
}

func NewClassTrackingService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	classTrackingRepository repositories.ClassTrackingRepository,
) ClassTrackingService {
	return &classTrackingService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:        redisCache,
		classTrackingRepository: classTrackingRepository,
	}
}

func (s *classTrackingService) InsertOrUpdate(ctx context.Context, request *dtos.InsertOrUpdateClassTrackingRequest) (*dtos.InsertOrUpdateClassTrackingResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	modelClassTracking := &loyalty.ClassTracking{

            
            Id: request.Id,

                        
            UserId: request.UserId,

            
            Source: request.Source,

            
            OldClassId: request.OldClassId,

            
            NewClassId: request.NewClassId,

            
            NumOfPoint: request.NumOfPoint,

            
            NumOfTrans: request.NumOfTrans,

            
            OldClassDate: request.OldClassDate,

            
            Created: request.Created,

            
            JsonDetail: request.JsonDetail,

            	}
	err := s.classTrackingRepository.InsertOrUpdate(ctx, tx, modelClassTracking)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "InsertOrUpdate classTracking error: %v", err)
	}
	resp := &dtos.InsertOrUpdateClassTrackingResponse{
		Meta: dtos.NewMetaOK(),
	}
	return resp, nil
}

func (s *classTrackingService) Delete(ctx context.Context, request *dtos.DeleteClassTrackingRequest) (*dtos.DeleteClassTrackingResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	err := s.classTrackingRepository.Delete(ctx, tx, request.Id)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Delete classTracking error: %v", err)
		return nil, err
	}
	resp := &dtos.DeleteClassTrackingResponse{
		Meta: dtos.NewMetaOK(),
	}
	return resp, nil
}

func (s *classTrackingService) GetByID(ctx context.Context, request *dtos.GetClassTrackingRequest) (*dtos.GetClassTrackingResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	item, err := s.classTrackingRepository.GetByID(ctx, tx, request.Id)
	if err != nil {
		logger.Errorf(ctx, "GetByID classTracking error: %v", err)
		return nil, err
	}
	data := &dtos.ClassTrackingInfo{}
	err = copier.Copy(&data, &item)
	if err != nil {
		logger.Errorf(ctx, "Copy classTracking error: %v", err)
		return nil, err
	}
	resp := &dtos.GetClassTrackingResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}

func (s *classTrackingService) List(ctx context.Context, request *dtos.ListClassTrackingRequest) (*dtos.ListClassTrackingResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	list, err := s.classTrackingRepository.List(ctx, tx, request.PageIndex, request.PageSize)
	if err != nil {
		logger.Errorf(ctx, "List classTracking error: %v", err)
		return nil, err
	}
	data := make([]*dtos.ClassTrackingInfo, 0)
	for _, item := range list {
		var (
			iv dtos.ClassTrackingInfo
			_  = copier.Copy(&iv, &item)
		)
		data = append(data, &iv)
	}
	resp := &dtos.ListClassTrackingResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}

