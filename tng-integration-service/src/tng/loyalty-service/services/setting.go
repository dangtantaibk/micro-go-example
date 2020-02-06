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

type SettingService interface {
	InsertOrUpdate(context.Context, *dtos.InsertOrUpdateSettingRequest) (*dtos.InsertOrUpdateSettingResponse, error)
	Delete(context.Context, *dtos.DeleteSettingRequest) (*dtos.DeleteSettingResponse, error)
	GetByID(context.Context, *dtos.GetSettingRequest) (*dtos.GetSettingResponse, error)
	List(context.Context, *dtos.ListSettingRequest) (*dtos.ListSettingResponse, error)
}

type settingService struct {
	BaseService
	redisCache        redisutil.Cache
	settingRepository repositories.SettingRepository
}

func NewSettingService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	settingRepository repositories.SettingRepository,
) SettingService {
	return &settingService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:        redisCache,
		settingRepository: settingRepository,
	}
}

func (s *settingService) InsertOrUpdate(ctx context.Context, request *dtos.InsertOrUpdateSettingRequest) (*dtos.InsertOrUpdateSettingResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	modelSetting := &loyalty.Setting{
		ID:                   request.ID,
		OutputMoneyPerPoint:  request.OutputMoneyPerPoint,
		PeriodOfClassByMonth: request.PeriodOfClassByMonth,
		JsonDetail:           request.JsonDetail,
	}
	err := s.settingRepository.InsertOrUpdate(ctx, tx, modelSetting)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "InsertOrUpdate setting error: %v", err)
	}
	resp := &dtos.InsertOrUpdateSettingResponse{
		Meta: dtos.NewMetaOK(),
	}
	return resp, nil
}

func (s *settingService) Delete(ctx context.Context, request *dtos.DeleteSettingRequest) (*dtos.DeleteSettingResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	err := s.settingRepository.Delete(ctx, tx, request.ID)
	if err != nil {
		logger.Errorf(ctx, "Delete setting error: %v", err)
		return nil, err
	}
	resp := &dtos.DeleteSettingResponse{
		Meta: dtos.NewMetaOK(),
	}
	return resp, nil
}

func (s *settingService) GetByID(ctx context.Context, request *dtos.GetSettingRequest) (*dtos.GetSettingResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	item, err := s.settingRepository.GetByID(ctx, tx, request.ID)
	if err != nil {
		logger.Errorf(ctx, "GetByID setting error: %v", err)
		return nil, err
	}
	data := &dtos.SettingInfo{}
	err = copier.Copy(&data, &item)
	if err != nil {
		logger.Errorf(ctx, "Copy setting error: %v", err)
		return nil, err
	}
	resp := &dtos.GetSettingResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}

func (s *settingService) List(ctx context.Context, request *dtos.ListSettingRequest) (*dtos.ListSettingResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	list, err := s.settingRepository.List(ctx, tx, request.PageIndex, request.PageSize)
	if err != nil {
		logger.Errorf(ctx, "List setting error: %v", err)
		return nil, err
	}
	data := make([]*dtos.SettingInfo, 0)
	for _, item := range list {
		var (
			iv dtos.SettingInfo
			_  = copier.Copy(&iv, &item)
		)
		data = append(data, &iv)
	}
	resp := &dtos.ListSettingResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}
