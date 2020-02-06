package services

import (
	"context"
	"fmt"
	beecontext "github.com/astaxie/beego/context"
	"github.com/jinzhu/copier"
	"tng/common/logger"
	"tng/common/models"
	models_shiper "tng/common/models/shiper"
	"tng/common/utils/db"
	"tng/common/utils/hashutil"
	jwtUtils "tng/common/utils/jwt"
	"tng/common/utils/redisutil"
	"tng/common/utils/strutil"
	"tng/shipper-service/dtos"
	"tng/shipper-service/repositories"
)

type shipperService struct {
	BaseService
	redisCache     redisutil.Cache
	shipperRepository repositories.ShipperRepository
}

func (s *shipperService) Delete(ctx context.Context, request *dtos.DeleteShipperAccountRequest) (*dtos.DeleteShipperAccountResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)

	defer s.dbFactory.Rollback(tx)
	err := s.shipperRepository.Delete(tx, request.ID)
	s.dbFactory.Commit(tx)

	if err != nil {
		logger.Errorf(ctx, "Delele shipper account error: %v", err)
		return nil, dtos.NewAppError(dtos.InternalServerError)
	}
	response := &dtos.DeleteShipperAccountResponse{Meta: dtos.Meta{
		Code:    1,
		Message: "OK",
	},}
	return response, nil
}

func (s *shipperService) List(ctx context.Context, request *dtos.ListShipperRequest) (*dtos.ListShipperResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)

	list, totalRecord, err := s.shipperRepository.List(tx, request.CurrentPage, request.TotalTransPerPage)

	if err != nil {
		logger.Errorf(ctx, "Get list shipper error: %v", err)
		return nil, err
	}

	data := make([]dtos.ShipperInfo, 0)
	for _, item := range list {
		var (
			iv dtos.ShipperInfo
			_  = copier.Copy(&iv, &item)
		)
		data = append(data, iv)
	}
	response := &dtos.ListShipperResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "ok",
		},
		Data: data,
		TotalRecord: totalRecord,
	}
	return response, nil
}

func (s *shipperService) Update(ctx context.Context, request *dtos.UpdateShipperRequest) (*dtos.UpdateShipperResponse, error) {
	var (
		tx            = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
		shipperModel = &models_shiper.Shipper{
			ID:       request.ID,
			Name:     request.Name,
			Phone:    request.Phone,
			Password: request.Password,
			Status:   request.Status,
		}
	)
	defer s.dbFactory.Rollback(tx)
	shipperInfo, err := s.shipperRepository.Update(tx, shipperModel)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Update item type error: %v", err)
		return nil, dtos.NewAppError(dtos.InternalServerError)
	}
	var shipperDto dtos.ShipperInfo
	err = copier.Copy(&shipperDto, &shipperInfo)
	if err != nil {
		logger.Errorf(ctx, "parse shipper type error: %v", err)
		return nil, err
	}

	response := &dtos.UpdateShipperResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
		Data: shipperDto,
	}
	return response, nil
}

func (s *shipperService) SignUp(ctx context.Context, request *dtos.SignUpShipperRequest) (*dtos.SignUpShipperResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)

	if strutil.IsEmpty(request.Password) || strutil.IsEmpty(request.Phone) {
		logger.Errorf(ctx, "Phone number or Password is empty")
		return nil, dtos.NewAppError(dtos.ErrorPasswordIncorrect)
	}

	password, err := hashutil.HashPassword(request.Password);
	if err != nil {
		logger.Errorf(ctx, "Hasing Password error: %v", err)
		return nil, dtos.NewAppError(dtos.ErrorUnknown)
	}

	modelInsertShipper := &models_shiper.Shipper{
		Status:    request.Status,
		Name:     request.Name,
		Phone: request.Phone,
		Password: password,
	}

	defer s.dbFactory.Rollback(tx)
	err = s.shipperRepository.Insert(tx, modelInsertShipper)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Error insert db: %v", err)
		return nil, dtos.NewAppError(dtos.InternalServerError)
	}
	response := &dtos.SignUpShipperResponse{
		SignUpSuccess: true,
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}
	return response, nil
}

func (s *shipperService) LoginWithPassword(ctx context.Context, request *dtos.LoginWithPasswordShipperRequest, input beecontext.BeegoInput) (*dtos.LoginWithPasswordShipperResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)
	shipperInfo, err := s.shipperRepository.GetByPhoneNumber(tx, request.PhoneNumber)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Phone number or Password is incorrect: %v", err)
		return nil, dtos.NewAppError(dtos.ErrorPasswordIncorrect)
	}
	if shipperInfo == nil {
		logger.Errorf(ctx, "Phone number or Password is incorrect")
		return nil, dtos.NewAppError(dtos.ErrorPasswordIncorrect)
	}

	match := hashutil.CheckPasswordHash(request.Password, shipperInfo.Password)
	if match == false {
		logger.Errorf(ctx, "Phone number or Password is incorrect")
		return nil, dtos.NewAppError(dtos.ErrorPasswordIncorrect)
	}

	tokenString, err := jwtUtils.GenerateShipperToken(shipperInfo, input)
	if err != nil {
		logger.Errorf(ctx, "Generate Token error: %v", err)
		return nil, err
	}

	response := &dtos.LoginWithPasswordShipperResponse{
		LoginSuccess: true,
		Token:        tokenString,
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}
	return response, nil
}

func (s *shipperService) VerifyPhoneNumber(ctx context.Context, request *dtos.VerifyPhoneNumberShipperRequest) (*dtos.VerifyPhoneNumberShipperResponse, error) {
	panic("implement me")
}

func (s *shipperService) RefreshToken(ctx context.Context, request *dtos.RefreshTokenShipperRequest) (*dtos.RefreshTokenShipperResponse, error) {
	panic("implement me")
}

type ShipperService interface {
	SignUp(ctx context.Context, request *dtos.SignUpShipperRequest) (*dtos.SignUpShipperResponse, error)
	LoginWithPassword(ctx context.Context, request *dtos.LoginWithPasswordShipperRequest, input beecontext.BeegoInput) (*dtos.LoginWithPasswordShipperResponse, error)
	VerifyPhoneNumber(ctx context.Context, request *dtos.VerifyPhoneNumberShipperRequest) (*dtos.VerifyPhoneNumberShipperResponse, error)
	RefreshToken(ctx context.Context, request *dtos.RefreshTokenShipperRequest) (*dtos.RefreshTokenShipperResponse, error)
	Update(ctx context.Context, request *dtos.UpdateShipperRequest) (*dtos.UpdateShipperResponse, error)
	List(ctx context.Context, request *dtos.ListShipperRequest) (*dtos.ListShipperResponse, error)
	Delete(ctx context.Context, request *dtos.DeleteShipperAccountRequest) (*dtos.DeleteShipperAccountResponse, error)
}

func NewShipperService(factory db.Factory,
	redisCache redisutil.Cache,
	shipperRepository repositories.ShipperRepository,) ShipperService {
	return &shipperService{
		BaseService: BaseService{
			dbFactory: factory,
		},
		redisCache:     redisCache,
		shipperRepository: shipperRepository,
	}

}