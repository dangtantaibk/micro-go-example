package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"tng/user-profile-service/dtos"
	"tng/user-profile-service/helper"
)

type UserSessionRepository interface {
	GetUserSession(context.Context, string) (*dtos.UserSessionInfo, error)
	CreateUserSession(context.Context, *dtos.UserSessionInfo) (*string, error)
	RemoveUserSession(context.Context, string) error
	UpdateUserSession(context.Context, *dtos.UpdateUserSessionToken) error
}

type userSessionRepository struct {
	helper helper.Helper
}

func NewUserSessionRepository(
	helper helper.Helper,
) UserSessionRepository {
	return &userSessionRepository{
		helper: helper,
	}
}

func (r *userSessionRepository) GetUserSession(ctx context.Context, token string) (*dtos.UserSessionInfo, error) {
	tokenInfo, err := r.helper.GetTokenInfo(token)
	if err != nil {
		return nil, err
	}

	userSession, err := r.helper.GetUserSession(tokenInfo)
	if err != nil {
		return nil, err
	}

	go r.helper.ValidateTTL(tokenInfo)

	return userSession, nil
}

func (r *userSessionRepository) CreateUserSession(ctx context.Context, request *dtos.UserSessionInfo) (*string, error) {
	userSession := request
	token, err := r.helper.GetSessionToken(userSession)
	if err != nil {
		return nil, err
	}

	tokenInfo := &helper.TokenInfo{
		UserId: fmt.Sprintf("%d", userSession.ID),
		Token:  token,
	}

	isTokenExist := len(tokenInfo.Token) > 0

	if !isTokenExist {
		tokenInfo.Token = r.helper.GenSessionToken()
		err = r.helper.StoreUserSession(tokenInfo, userSession)
		if err != nil {
			return nil, err
		}
	}

	// Generate and response the new Token normally!
	newToken, err := r.helper.GenToken(tokenInfo)
	if err != nil {
		return nil, err
	}

	// Additional, if User Session already existed. Need to upgrade it before!
	if isTokenExist {
		updateUserSessionReq := &dtos.UpdateUserSessionToken{
			Token:       newToken,
			UserSession: userSession,
		}
		err = r.UpdateUserSession(ctx, updateUserSessionReq)
		if err != nil {
			return nil, err
		}
	}

	return &newToken, nil
}

func (r *userSessionRepository) RemoveUserSession(ctx context.Context, token string) error {
	tokenInfo, err := r.helper.GetTokenInfo(token)
	if err != nil {
		return err
	}

	err = r.helper.RemoveUserSession(tokenInfo)
	if err != nil {
		return err
	}

	return nil
}

func (r *userSessionRepository) UpdateUserSession(ctx context.Context, request *dtos.UpdateUserSessionToken) error {
	token := request.Token
	userSession := request.UserSession

	var tokenInfo *helper.TokenInfo
	var err error

	//TH: cập nhật us từ middleware
	// 1. có public token
	// 2. không có public token, chỉ có user_id
	if len(token) > 0 {
		tokenInfo, err = r.helper.GetTokenInfo(token)
		if err != nil {
			return err
		}
	} else {
		tokenInfo = &helper.TokenInfo{
			UserId: fmt.Sprintf("%d", userSession.ID),
		}
	}

	// áp dụng logic merge us tương tự luồng từ kafka
	newUserSession, err := r.mergeUserSession(tokenInfo, userSession)
	if err != nil {
		return err
	}

	err = r.helper.UpdateUserSession(tokenInfo, newUserSession)
	if err != nil {
		return err
	}

	return nil
}

func (s *userSessionRepository) mergeUserSession(tokenInfo *helper.TokenInfo, us *dtos.UserSessionInfo) (*dtos.UserSessionInfo, error) {
	val, err := json.Marshal(us)
	if err != nil {
		return nil, err
	}

	userSession, err := s.helper.GetUserSession(tokenInfo)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(val, &userSession)
	if err != nil {
		return nil, err
	}
	return userSession, nil
}
