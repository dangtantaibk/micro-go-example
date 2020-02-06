package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"tng/common/concurrency"
	"tng/common/location"
	"tng/common/logger"
	"tng/common/models"
	userProfile "tng/common/models/user-profile"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/user-profile-service/dtos"
	"tng/user-profile-service/helper"
	"tng/user-profile-service/repositories"
)

type UserProfileService interface {
	Create(context.Context, *dtos.CreateProfileRequest) (*dtos.CreateProfileResponse, error)
	Update(context.Context, *dtos.UpdateProfileRequest) (*dtos.UpdateProfileResponse, error)
	GetByID(context.Context, *dtos.GetProfileByIDRequest) (*dtos.GetProfileResponse, error)
}

type userProfileService struct {
	BaseService
	redisCache     redisutil.Cache
	userRepository repositories.UserProfileRepository
}

func NewUserProfileService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	userRepository repositories.UserProfileRepository,
) UserProfileService {
	return &userProfileService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:     redisCache,
		userRepository: userRepository,
	}
}
func (s *userProfileService) Create(ctx context.Context, request *dtos.CreateProfileRequest) (*dtos.CreateProfileResponse, error) {
	dataRequest := request.Data
	if dataRequest == nil {
		return nil, fmt.Errorf("data request nil")
	}

	count, err := s.setMysqlUserProfile(dataRequest)
	if err != nil {
		logger.Errorf(ctx, "insert mysql error: %v", err)
		return nil, dtos.NewAppError(dtos.InsertUserProfileError)
	}
	if count <= 0 {
		logger.Errorf(ctx, "count insert mysql error: %d", count)
		return nil, dtos.NewAppError(dtos.InsertUserProfileError)
	} else {
		go func() {
			dataRequest.ID = count
			s.setRedisUserProfile(dataRequest)
		}()
	}
	response := &dtos.CreateProfileResponse{
		Meta: dtos.MetaOK(),
	}
	return response, nil
}

func (s *userProfileService) Update(ctx context.Context, request *dtos.UpdateProfileRequest) (*dtos.UpdateProfileResponse, error) {
	dataRequest := request.Data
	if dataRequest == nil {
		return nil, fmt.Errorf("data request nil")
	}
	cc := concurrency.New()
	cc.Add(func() error {
		return s.setRedisUserProfile(dataRequest)
	})
	cc.Add(func() error {
		return s.updateMysqlUserProfile(dataRequest)
	})
	if err := cc.Do(); err != nil {
		logger.Errorf(ctx, "update user profile error: %v", err)
		return nil, dtos.NewAppError(dtos.UpdateUserProfileError)
	}
	response := &dtos.UpdateProfileResponse{
		Meta: dtos.MetaOK(),
	}
	return response, nil
}

func (s *userProfileService) GetByID(ctx context.Context, request *dtos.GetProfileByIDRequest) (*dtos.GetProfileResponse, error) {
	dataUser, err := s.getUserProfileRedisByUserID(request)
	if err == nil && dataUser != nil {
		return &dtos.GetProfileResponse{
			Data: dataUser,
			Meta: dtos.MetaOK(),
		}, nil
	}
	dataUser, err = s.getUserProfileMySqlByUserID(request)
	if err != nil {
		logger.Errorf(ctx, "get user profile error: %v", err)
		return nil, dtos.NewAppError(dtos.GetUserProfileError)
	}

	return &dtos.GetProfileResponse{
		Data: dataUser,
		Meta: dtos.MetaOK(),
	}, nil
}

func (s *userProfileService) setRedisUserProfile(data *dtos.User) error {
	k := helper.RedisKeyUserProfile(fmt.Sprintf("%d", data.ID), fmt.Sprintf("%d", data.AppID))
	return s.redisCache.Set(k, data, 0)
}

func (s *userProfileService) setMysqlUserProfile(data *dtos.User) (int64, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	userModel := &userProfile.User{
		AppID:      data.AppID,
		UCode:      data.UCode,
		Title:      data.Title,
		FirstName:  data.FirstName,
		SurName:    data.SurName,
		FullName:   data.FullName,
		LastName:   data.LastName,
		Phone:      data.Phone,
		HomePhone:  data.HomePhone,
		Email:      data.Email,
		Address:    data.Address,
		SocialID:   data.SocialID,
		LoginType:  data.LoginType,
		WardID:     data.WardID,
		DistrictID: data.DistrictID,
		ProvinceID: data.ProvinceID,
		CountryID:  data.CountryID,
		Created:    location.GetVNCurrentTime().Format(models.FormatYYYMMDDHHMMSS),
		CreatedBy:  data.CreatedBy,
		Avatar:     data.Avatar,
		Status:     data.Status,
		Lat:        data.Lat,
		Long:       data.Long,
	}
	count, err := s.userRepository.Create(tx, userModel)
	if err != nil {
		return 0, err
	}
	s.dbFactory.Commit(tx)
	return count, nil
}

func (s *userProfileService) updateMysqlUserProfile(data *dtos.User) error {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	userModel := &userProfile.User{
		ID:         data.ID,
		UCode:      data.UCode,
		Title:      data.Title,
		FirstName:  data.FirstName,
		SurName:    data.SurName,
		FullName:   data.FullName,
		LastName:   data.LastName,
		Phone:      data.Phone,
		HomePhone:  data.HomePhone,
		Email:      data.Email,
		Address:    data.Address,
		SocialID:   data.SocialID,
		LoginType:  data.LoginType,
		WardID:     data.WardID,
		DistrictID: data.DistrictID,
		ProvinceID: data.ProvinceID,
		CountryID:  data.CountryID,
		Avatar:     data.Avatar,
		Status:     data.Status,
		Lat:        data.Lat,
		Long:       data.Long,
	}
	err := s.userRepository.Update(tx, userModel)
	if err != nil {
		return err
	}
	s.dbFactory.Commit(tx)
	return nil
}

func (s *userProfileService) getUserProfileRedisByUserID(request *dtos.GetProfileByIDRequest) (*dtos.User, error) {
	k := helper.RedisKeyUserProfile(fmt.Sprintf("%d", request.UserID), fmt.Sprintf("%d", request.AppID))
	val, err := s.redisCache.Get(k)
	if err != nil {
		return nil, err
	}
	dataUser := &dtos.User{}
	err = json.Unmarshal([]byte(val), dataUser)
	if err != nil {
		return nil, err
	}
	return dataUser, nil
}

func (s *userProfileService) getUserProfileMySqlByUserID(request *dtos.GetProfileByIDRequest) (*dtos.User, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)

	profile, err := s.userRepository.GetByID(tx, fmt.Sprintf("%d", request.UserID))
	if err != nil {
		return nil, err
	}
	dataUser := &dtos.User{}
	err = copier.Copy(&dataUser, &profile)
	if err != nil {
		return nil, err
	}
	if dataUser != nil {
		go s.setRedisUserProfile(dataUser)
	}
	return dataUser, nil
}
