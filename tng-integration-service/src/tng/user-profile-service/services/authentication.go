package services

import (
	"context"
	"github.com/jinzhu/copier"
	zaloSocialApi "tng/common/adapters/zalo"
	"tng/common/logger"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/user-profile-service/dtos"
	"tng/user-profile-service/helper"
	"tng/user-profile-service/repositories"
)

type AuthenticationService interface {
	Login(context.Context, *dtos.DataLoginRequest) (*dtos.LoginResponse, error)
	CheckLogin(context.Context, *dtos.CheckLoginRequest) (*dtos.CheckLoginResponse, error)
}

type authenticationService struct {
	BaseService
	redisCache            redisutil.Cache
	userSessionRepository repositories.UserSessionRepository
	userProfileRepository repositories.UserProfileRepository
	zaloLoginApi          zaloSocialApi.Adapter
	helper                helper.Helper
}

func NewAuthenticationService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	userSessionRepository repositories.UserSessionRepository,
	userProfileRepository repositories.UserProfileRepository,
	zaloLoginApi zaloSocialApi.Adapter,
	helper helper.Helper,
) AuthenticationService {
	return &authenticationService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:            redisCache,
		userSessionRepository: userSessionRepository,
		userProfileRepository: userProfileRepository,
		zaloLoginApi:          zaloLoginApi,
		helper:                helper,
	}
}

func (s *authenticationService) Login(ctx context.Context, dataLogin *dtos.DataLoginRequest) (*dtos.LoginResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	profile, err := s.userProfileRepository.GetBySocialID(tx, dataLogin.SocialID)
	if err != nil {
		return nil, dtos.NewAppError(dtos.LoginInfoInvalid)
	}
	userSession := &dtos.UserSessionInfo{}
	err = copier.Copy(&userSession, &profile)
	if err != nil {
		logger.Errorf(ctx, "copy user profile err: %v", err)
		return nil, dtos.NewAppError(dtos.LoginInfoInvalid)
	}
	token, err := s.userSessionRepository.CreateUserSession(ctx, userSession)
	data := &dtos.LoginResponseData{
		Token: *token,
	}
	response := &dtos.LoginResponse{
		Data: data,
		Meta: dtos.MetaOK(),
	}
	return response, nil
}

func (s *authenticationService) CheckLogin(ctx context.Context, request *dtos.CheckLoginRequest) (*dtos.CheckLoginResponse, error) {
	userData, err := s.userSessionRepository.GetUserSession(ctx, request.Token)
	if err != nil {
		logger.Errorf(ctx, "get user profile from token error: %v", err)
		return nil, dtos.NewAppError(dtos.LoginTokenInvalid)
	}
	if userData == nil {
		logger.Errorf(ctx, "get user profile from token error")
		return nil, dtos.NewAppError(dtos.LoginTokenInvalid)
	}
	return &dtos.CheckLoginResponse{
		Meta: dtos.MetaOK(),
	}, nil
}
