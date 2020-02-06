package services

import (
	"context"
	"tng/common/logger"
	"tng/common/models/dev-tool"
	"tng/common/utils/db"
	"tng/dev-tool-service/dtos"
	"tng/dev-tool-service/repositories"
)

// Pi3UpdateInfoService represents a service for update info.
type Pi3UpdateInfoService interface {
	GetPi3UpdateInfo(ctx context.Context, request *dtos.GetPi3UpdateInfoRequest) (*dtos.GetPi3UpdateInfoResponse, error)
	RegisterPi3(ctx context.Context, request *dtos.RegisterPi3Request) (*dtos.RegisterPi3Response, error)
}

type pi3UpdateInfoService struct {
	BaseService
	pi3UpdateInfoRepository repositories.Pi3UpdateInfoRepository
}

// NewPi3UpdateInfoService create a new instance for update info
func NewPi3UpdateInfoService(
	dbFactory db.Factory,
	pi3UpdateInfoRepository repositories.Pi3UpdateInfoRepository,
) Pi3UpdateInfoService {
	return &pi3UpdateInfoService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		pi3UpdateInfoRepository: pi3UpdateInfoRepository,
	}
}

func (s *pi3UpdateInfoService) GetPi3UpdateInfo(ctx context.Context, request *dtos.GetPi3UpdateInfoRequest) (*dtos.GetPi3UpdateInfoResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	resp, err := s.pi3UpdateInfoRepository.GetPi3UpdateInfo(tx, request.PosID)
	if err != nil {
		logger.Errorf(ctx, "get pi3 update info error: %v", err)
		return nil, err
	}
	if resp == nil {
		logger.Errorf(ctx, "resp nil")
		return nil, dtos.NewAppError(dtos.ErrorNotFoundData)
	}
	return &dtos.GetPi3UpdateInfoResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
		Data: &dtos.DataPi3UpdateInfo{
			URL: resp.UrlFileExecute,
		},
	}, nil
}

func (s *pi3UpdateInfoService) RegisterPi3(ctx context.Context, request *dtos.RegisterPi3Request) (*dtos.RegisterPi3Response, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	info := &dev_tool.DeviceUpdateInfo{
		PosID:      request.PosID,
		DeviceType: request.DeviceType,
	}

	err := s.pi3UpdateInfoRepository.SetPi3UpdateInfo(tx, info)
	if err != nil {
		logger.Errorf(ctx, "register device error: %v", err)
		return nil, dtos.NewAppError(dtos.ErrorRegisterDevice)
	}
	return &dtos.RegisterPi3Response{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}, nil
}
