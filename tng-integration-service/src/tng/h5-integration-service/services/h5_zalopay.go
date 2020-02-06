package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	_h5ZaloPayAdapter "tng/common/adapters/h5-zalopay"
	"tng/common/concurrency"
	"tng/common/logger"
	"tng/common/models/h5-intergration"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/h5-integration-service/dtos"
	"tng/h5-integration-service/repositories"
)

// H5ZaloPayService represents a service for managing H5ZaloPay.
type H5ZaloPayService interface {
	GrantedMBToken(ctx context.Context, request *dtos.GetH5ZaloPayRequest) (*dtos.GetH5ZaloPayResp, error)
	GetPaymentOrderUrl(ctx context.Context, request *dtos.H5ZaloPayOrderURLReq) (*dtos.H5ZaloPayOrderURLResp, error)
}

type h5ZaloPayService struct {
	BaseService
	redisCache          redisutil.Cache
	h5ZaloPayAdapter    _h5ZaloPayAdapter.Adapter
	h5ZaloPayRepository repositories.H5ZalopayRepository
}

// NewH5ZaloPayService create a new instance for H5ZaloPay
func NewH5ZaloPayService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	h5ZaloPayAdapter _h5ZaloPayAdapter.Adapter,
	h5ZaloPayRepository repositories.H5ZalopayRepository,
) H5ZaloPayService {
	return &h5ZaloPayService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:          redisCache,
		h5ZaloPayAdapter:    h5ZaloPayAdapter,
		h5ZaloPayRepository: h5ZaloPayRepository,
	}
}

func (h *h5ZaloPayService) GrantedMBToken(ctx context.Context, req *dtos.GetH5ZaloPayRequest) (*dtos.GetH5ZaloPayResp, error) {
	cc := concurrency.New()
	cc.Add(func() error {
		externalResp, err := h.h5ZaloPayAdapter.GetMBToken(ctx, fmt.Sprintf("%d", req.AppID), req.MAToken)
		if err != nil {
			return dtos.NewAppError(dtos.ErrorUnknown)
		}
		if externalResp == nil {
			return dtos.NewAppError(dtos.ErrorDataOfH5Incorrect)
		}
		if externalResp.ReturnCode != 1 {
			return dtos.NewAppError(dtos.ErrorDataOfH5Incorrect)
		}
		if externalResp.Data == nil {
			return dtos.NewAppError(dtos.ErrorDataOfH5Incorrect)
		}
		mbToken := externalResp.Data.MBToken
		if mbToken == "" {
			logger.Errorf(ctx, "matoken nil")
			return dtos.NewAppError(dtos.ErrorDataOfH5Incorrect)
		} else {
			go func() {
				tkInfo := &dtos.TokenInfo{
					AppID:      req.AppID,
					TokenType:  dtos.TOKEN_TYPE_H5_ZALOPAY,
					MAToken:    req.MAToken,
					MBToken:    mbToken,
					ExpireTime: 0,
					Status:     dtos.TOKEN_STT_ENABLE,
					UserID:     req.UserID,
				}
				err := h.setData(ctx, *tkInfo)
				if err != nil {
					logger.Errorf(ctx, "set data error: %v", err)
				}
			}()
		}
		return nil
	})
	if err := cc.Do(); err != nil {
		logger.Errorf(ctx, "do cc error: %v", err)
		return nil, err
	}

	resp := &dtos.GetH5ZaloPayResp{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}
	return resp, nil
}

func (h *h5ZaloPayService) GetPaymentOrderUrl(ctx context.Context, req *dtos.H5ZaloPayOrderURLReq) (*dtos.H5ZaloPayOrderURLResp, error) {
	tkInfo, err := h.getData(ctx, req.AppID, req.UserID)
	if err != nil {
		logger.Errorf(ctx, "get token info error: %v", err)
		return nil, dtos.NewAppError(dtos.ErrorNotFoundData)
	}
	if tkInfo == nil {
		logger.Errorf(ctx, "token info nil")
		return nil, dtos.NewAppError(dtos.ErrorNotFoundData)
	}
	mbToken := tkInfo.MBToken
	maToken := tkInfo.MAToken
	if maToken != req.MAToken {
		logger.Errorf(ctx, "matoken not matching")
		return nil, dtos.NewAppError(dtos.ErrorMBTokenIsChange)
	}
	if tkInfo.Status != dtos.TOKEN_STT_ENABLE {
		logger.Errorf(ctx, "account is blocked")
		return nil, dtos.NewAppError(dtos.ErrorMBTokenIsChange)
	}
	externalReq := &_h5ZaloPayAdapter.H5ZaloPayOrderURLReq{
		ZPOrderURL: req.ZPOrderURL,
		AppID:      req.AppID,
		MBToken:    mbToken,
	}
	externalResp, err := h.h5ZaloPayAdapter.GetOrder(ctx, externalReq)
	if err != nil {
		return nil, dtos.NewAppError(dtos.ErrorUnknown)
	}
	if externalResp == nil {
		return nil, dtos.NewAppError(dtos.ErrorDataOfH5Incorrect)
	}
	if externalResp.Data == nil {
		logger.Errorf(ctx, "data of order url nil")
		return nil, dtos.NewAppError(dtos.ErrorDataOfH5Incorrect)
	}
	paymentUrl := externalResp.Data.PaymentURL
	if paymentUrl == "" {
		logger.Errorf(ctx, "payment url nil")
		return nil, dtos.NewAppError(dtos.ErrorDataOfH5Incorrect)
	}
	data := &dtos.DataH5ZaloPayOrderURL{
		PaymentURL: paymentUrl,
	}
	resp := &dtos.H5ZaloPayOrderURLResp{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
		Data: data,
	}
	return resp, nil
}

func (h *h5ZaloPayService) getData(ctx context.Context, appID int, userID string) (*dtos.TokenInfo, error) {
	tokenInfo, err := h.getDataInRedis(ctx, appID, userID)
	if err == nil || tokenInfo != nil {
		return tokenInfo, nil
	}
	tokenInfo, err = h.getDataInMySql(ctx, appID, userID)
	if err == nil || tokenInfo != nil {
		go func() {
			k := fmt.Sprintf("%d-%s", appID, userID)
			v, err := json.Marshal(tokenInfo)
			if err != nil {
				return
			}
			err = h.redisCache.Set(k, string(v), 0)
			if err != nil {
				return
			}
		}()
		return tokenInfo, nil
	}
	return nil, fmt.Errorf("get data err")
}

func (h *h5ZaloPayService) getDataInRedis(ctx context.Context, appID int, userID string) (*dtos.TokenInfo, error) {
	key := fmt.Sprintf("%d-%s", appID, userID)
	tkInfo := &dtos.TokenInfo{}
	result, err := h.redisCache.Get(key)
	if err != nil {
		return nil, err
	}
	jsonInput, err := strconv.Unquote(string(result))
	err = json.Unmarshal([]byte(jsonInput), tkInfo)
	if err != nil {
		return nil, err
	}
	maToken := tkInfo.MAToken
	if maToken == "" {
		return nil, fmt.Errorf("matoken nil")
	}
	return tkInfo, nil
}

func (h *h5ZaloPayService) getDataInMySql(ctx context.Context, appID int, userID string) (*dtos.TokenInfo, error) {
	var (
		tx = h.dbFactory.GetDB(true)
	)
	resp, err := h.h5ZaloPayRepository.GetTokenInfo(tx, appID, userID)
	h.dbFactory.Commit(tx)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("token info nil")
	}
	tkInfo := &dtos.TokenInfo{
		AppID:      resp.AppID,
		TokenType:  resp.TokenType,
		MAToken:    resp.MAToken,
		MBToken:    resp.MBToken,
		ExpireTime: resp.ExpireTime,
		Status:     resp.Status,
		UserID:     resp.UserID,
	}
	return tkInfo, nil
}

// setData to sql and cache
func (h *h5ZaloPayService) setData(ctx context.Context, dataInput dtos.TokenInfo) error {
	cc := concurrency.New()
	// set data to redis
	cc.Add(func() error {
		k := fmt.Sprintf("%d-%s", dataInput.AppID, dataInput.UserID)
		v, err := json.Marshal(dataInput)
		if err != nil {
			return err
		}
		err = h.redisCache.Set(k, string(v), 0)
		if err != nil {
			return err
		}
		return nil
	})
	// set data to mysql
	cc.Add(func() error {
		var (
			tx = h.dbFactory.GetDB(true)
		)
		defer h.dbFactory.Rollback(tx)
		modelToken := h5_intergration.H5ZaloPay{
			AppID:      dataInput.AppID,
			TokenType:  dataInput.TokenType,
			MAToken:    dataInput.MAToken,
			MBToken:    dataInput.MBToken,
			Status:     dataInput.Status,
			ExpireTime: dataInput.ExpireTime,
			UserID:     dataInput.UserID,
		}
		err := h.h5ZaloPayRepository.SetTokenInfo(tx, &modelToken)
		h.dbFactory.Commit(tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err := cc.Do(); err != nil {
		return fmt.Errorf("do cc error when set data: %v", err)
	}
	return nil
}
