package services

import (
	"context"
	"fmt"
	beecontext "github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"time"
	_zaloSocialApi "tng/common/adapters/zalo"
	"tng/common/logger"
	"tng/common/models"
	models_shiper "tng/common/models/shiper"
	"tng/common/utils/cfgutil"
	"tng/common/utils/db"
	"tng/common/utils/hashutil"
	jwtUtils "tng/common/utils/jwt"
	"tng/common/utils/redisutil"
	"tng/common/utils/strutil"
	"tng/shipper-service/dtos"
	"tng/shipper-service/repositories"
)

type UserService interface {
	LoginWithZalo(ctx context.Context, request *dtos.LoginRequest, input beecontext.BeegoInput) (*dtos.LoginResponse, error)
	SignUp(ctx context.Context, request *dtos.SignUpRequest) (*dtos.SignUpResponse, error)
	LoginWithPassword(ctx context.Context, request *dtos.LoginWithPasswordRequest, input beecontext.BeegoInput) (*dtos.LoginWithPasswordResponse, error)
	VerifyPhoneNumber(ctx context.Context, request *dtos.VerifyPhoneNumberRequest) (*dtos.VerifyPhoneNumberResponse, error)
	RefreshToken(ctx context.Context, request *dtos.RefreshTokenRequest) (*dtos.RefreshTokenResponse, error)
}

type userService struct {
	BaseService
	redisCache     redisutil.Cache
	userRepository repositories.UserRepository
	zaloLoginApi   _zaloSocialApi.ZaloLoginApi
}



// NewUserService create a new instance
func NewUserService(factory db.Factory,
	redisCache redisutil.Cache,
	userRepository repositories.UserRepository,
	zaloLoginApi _zaloSocialApi.ZaloLoginApi, ) UserService {
	return &userService{
		BaseService: BaseService{
			dbFactory: factory,
		},
		redisCache:     redisCache,
		userRepository: userRepository,
		zaloLoginApi:   zaloLoginApi,
	}

}
func (s *userService) RefreshToken(ctx context.Context, request *dtos.RefreshTokenRequest) (*dtos.RefreshTokenResponse, error) {

	if strutil.IsEmpty(request.OldToken) {
		logger.Errorf(ctx, "Old token is empty")
		return nil, dtos.NewAppError(dtos.InvalidRequestError)
	}
	oldToken := request.OldToken
	claims := &dtos.Claims{}

	tkn, err := jwt.ParseWithClaims(oldToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfgutil.Load("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			logger.Errorf(ctx, "signature is invalid")
			return nil, dtos.NewAppError(dtos.UnauthorizedError)
		}
		logger.Errorf(ctx, "Unknown Error: %v", err)
		return nil, dtos.NewAppError(dtos.UnknownError)
	}
	if !tkn.Valid {
		logger.Errorf(ctx, "token is invalid")
		return nil, dtos.NewAppError(dtos.UnauthorizedError)
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 0 {
		logger.Errorf(ctx, "token is invalid")
		return nil, dtos.NewAppError(dtos.UnauthorizedError)
	}


	expirationTime := time.Now().Add(15 * 24 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString([]byte(cfgutil.Load("JWT_SECRET_KEY")))
	if err != nil {
		logger.Errorf(ctx, "Generate Token Error: %v", err)
		return nil, dtos.NewAppError(dtos.UnknownError)
	}

	response := &dtos.RefreshTokenResponse{
		NewToken: tokenString,
		Meta:     dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}
	return response, nil
}

func (s *userService) VerifyPhoneNumber(ctx context.Context, request *dtos.VerifyPhoneNumberRequest) (*dtos.VerifyPhoneNumberResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)

	userInfo, err := s.userRepository.GetByPhoneNumber(tx, request.PhoneNumber)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Error get phone number: %v", err)
		return nil, dtos.NewAppError(dtos.ErrorPhoneNumberNotFound)
	}
	if userInfo == nil {
		logger.Errorf(ctx, "Phone number is not exits")
		return nil, dtos.NewAppError(dtos.ErrorPhoneNumberNotFound)
	}

	response := &dtos.VerifyPhoneNumberResponse{
		VerifySuccess: true,
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}
	return response, nil
}

func (s *userService) SignUp(ctx context.Context, request *dtos.SignUpRequest) (*dtos.SignUpResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)

	if strutil.IsEmpty(request.Password) || strutil.IsEmpty(request.PhoneNumber) {
		logger.Errorf(ctx, "Phone number or Password is empty")
		return nil, dtos.NewAppError(dtos.ErrorPasswordIncorrect)
	}

	password, err := hashutil.HashPassword(request.Password);
	if err != nil {
		logger.Errorf(ctx, "Hasing Password error: %v", err)
		return nil, dtos.NewAppError(dtos.ErrorUnknown)
	}

	modelInsertUser := &models_shiper.User{
		SocialName:  "",
		SocialId:    0,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		Password: password,
	}

	defer s.dbFactory.Rollback(tx)
	err = s.userRepository.Insert(tx, modelInsertUser)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Error insert db: %v", err)
		return nil, dtos.NewAppError(dtos.InternalServerError)
	}
	response := &dtos.SignUpResponse{
		SignUpSuccess: true,
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}
	return response, nil
}

func (s *userService) LoginWithZalo(ctx context.Context, request *dtos.LoginRequest, input beecontext.BeegoInput) (*dtos.LoginResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)
	modelInsertUser := &models_shiper.User{
		SocialName:  "zalo",
		SocialId:    request.UserId,
		Address:     "",
		PhoneNumber: "",
	}

	getAccessTokenDto := &_zaloSocialApi.ZaloGetAccessTokenRequest{
		AppID:       request.AppID,
		AppSecret:   "",
		Code:        request.Code,
		RedirectUri: request.RedirectURI,
		IsSDK:       "true",
	}

	accessToken, errGetToken := s.zaloLoginApi.GetAccessToken(getAccessTokenDto)

	if errGetToken != nil {
		logger.Errorf(ctx, "Get Access token error: %v", errGetToken)
		return nil, dtos.NewAppError(dtos.UnauthorizedError)
	}

	if accessToken == nil {
		logger.Errorf(ctx, "Access token nil")
		return nil, dtos.NewAppError(dtos.UnauthorizedError)
	}

	if accessToken.AccessToken == "" {
		logger.Errorf(ctx, "Access token empty")
		return nil, dtos.NewAppError(dtos.UnauthorizedError)
	}

	defer s.dbFactory.Rollback(tx)
	err := s.userRepository.Insert(tx, modelInsertUser)
	s.dbFactory.Commit(tx)

	if err != nil {
		logger.Errorf(ctx, "Service insert error: %v", err)
		return nil, err
	}

	tokenString, err := jwtUtils.GenerateToken(modelInsertUser, input)
	if err != nil {
		logger.Errorf(ctx, "Generate Token error: %v", err)
		return nil, err
	}

	response := &dtos.LoginResponse{
		LoginSuccess: true,
		Token:        tokenString,
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}
	return response, nil
}

func (s *userService) LoginWithPassword(ctx context.Context, request *dtos.LoginWithPasswordRequest, input beecontext.BeegoInput) (*dtos.LoginWithPasswordResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)
	userInfo, err := s.userRepository.GetByPhoneNumber(tx, request.PhoneNumber)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Phone number or Password is incorrect: %v", err)
		return nil, dtos.NewAppError(dtos.ErrorPasswordIncorrect)
	}
	if userInfo == nil {
		logger.Errorf(ctx, "Phone number or Password is incorrect")
		return nil, dtos.NewAppError(dtos.ErrorPasswordIncorrect)
	}

	match := hashutil.CheckPasswordHash(request.Password, userInfo.Password)
	if match == false {
		logger.Errorf(ctx, "Phone number or Password is incorrect")
		return nil, dtos.NewAppError(dtos.ErrorPasswordIncorrect)
	}
	user := &models_shiper.User{
		SocialName:  "zalo",
		PhoneNumber: userInfo.Phone,
	}

	tokenString, err := jwtUtils.GenerateToken(user, input)
	if err != nil {
		logger.Errorf(ctx, "Generate Token error: %v", err)
		return nil, err
	}

	response := &dtos.LoginWithPasswordResponse{
		LoginSuccess: true,
		Token:        tokenString,
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}
	return response, nil


}


