package services

import (
	"context"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/promotion-service/dtos"
	"tng/promotion-service/repositories"
)

// PromotionService represents a service.
type PromotionService interface {
	Insert(context.Context, *dtos.PromotionRequest) (*dtos.PromotionResponse, error)
}

type promotionService struct {
	BaseService
	redisCache          redisutil.Cache
	promotionRepository repositories.PromotionRepository
}

// NewCampaignService create a new instance
func NewPromotionService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	promotionRepository repositories.PromotionRepository) PromotionService {
	return &promotionService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:          redisCache,
		promotionRepository: promotionRepository,
	}
}

func (s *promotionService) Insert(ctx context.Context, request *dtos.PromotionRequest) (*dtos.PromotionResponse, error) {
	resp := &dtos.PromotionResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}
	return resp, nil
}
