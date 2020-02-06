package services

import (
	"context"
	"github.com/jinzhu/copier"
	"time"
	"tng/common/logger"
	model_promotion "tng/common/models/promotion"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/promotion-service/dtos"
	"tng/promotion-service/repositories"
)

// CampaignService represents a service.
type CampaignService interface {
	Insert(context.Context, *dtos.InsertCampaignRequest) (*dtos.InsertCampaignResponse, error)
	List(context.Context, *dtos.ListCampaignRequest) (*dtos.ListCampaignResponse, error)
}

type campaignService struct {
	BaseService
	redisCache         redisutil.Cache
	campaignRepository repositories.CampaignRepository
}

// NewCampaignService create a new instance
func NewCampaignService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	campaignRepository repositories.CampaignRepository,
) CampaignService {
	return &campaignService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:         redisCache,
		campaignRepository: campaignRepository,
	}
}

func (s *campaignService) Insert(ctx context.Context, request *dtos.InsertCampaignRequest) (*dtos.InsertCampaignResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	modelInsertCampaign := &model_promotion.Campaign{
		Title:           request.Title,
		From:            request.From,
		To:              request.To,
		RepeatType:      request.RepeatType,
		RepeateValue:    request.RepeateValue,
		ExcludeDays:     request.ExcludeDays,
		Status:          request.Status,
		TimeFrom:        request.TimeFrom,
		TimeTo:          request.TimeTo,
		AppID:           request.AppID,
		MinOrderAmt:     request.MinOrderAmt,
		MaxOrderAmt:     request.MaxOrderAmt,
		MaxAmt:          request.MaxAmt,
		MaxTransaction:  request.MaxTransaction,
		PromoType:       request.PromoType,
		SchemeType:      request.SchemeType,
		DiscountPercent: request.DiscountPercent,
		DiscountAmt:     request.DiscountAmt,
		FixedDiscount:   request.FixedDiscount,
		Created:         time.Now(),
		CreatedBy:       "",
		Modified:        time.Now(),
		ModifiedBy:      "",
		Description:     request.Description,
		Json:            request.Json,
		UCode:           request.UCode,
	}
	defer s.dbFactory.Rollback(tx)
	err := s.campaignRepository.Insert(tx, modelInsertCampaign)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Service insert error: %v", err)
		return nil, err
	}

	resp := &dtos.InsertCampaignResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}
	return resp, nil
}

func (s *campaignService) List(ctx context.Context, request *dtos.ListCampaignRequest) (*dtos.ListCampaignResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	campaigns, err := s.campaignRepository.List(tx, request)
	if err != nil {
		logger.Errorf(ctx, "Get list campaign error: %v", err)
		return nil, err
	}
	data := make([]*dtos.Campaign, 0)
	for _, it := range campaigns {
		var (
			iv dtos.Campaign
			_  = copier.Copy(&iv, &it)
		)
		data = append(data, &iv)
	}
	resp := &dtos.ListCampaignResponse{
		Meta: dtos.MetaOK(),
		Data: data,
	}
	return resp, nil
}
