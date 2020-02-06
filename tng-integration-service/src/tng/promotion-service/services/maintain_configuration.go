package services

import (
	"context"
	"github.com/astaxie/beego/orm"
	"tng/common/logger"
	"tng/common/models"
	"tng/common/utils/db"
	"tng/promotion-service/repositories"
)

// MaintainConfigurationService contains all business of MaintainConfiguration.
type MaintainConfigurationService interface {
	IsDisable(ctx context.Context, module, action string) bool
}

type maintainConfigurationService struct {
	BaseService
	maintainConfigurationRepository repositories.MaintainConfigurationRepository
}

// NewMaintainConfigurationService return a new instance of MaintainConfigurationService.
func NewMaintainConfigurationService(
	dbFactory db.Factory,
	maintainConfigurationRepository repositories.MaintainConfigurationRepository,
) MaintainConfigurationService {
	s := &maintainConfigurationService{
		BaseService:                     BaseService{dbFactory: dbFactory},
		maintainConfigurationRepository: maintainConfigurationRepository,
	}
	return s
}

func (s *maintainConfigurationService) IsDisable(ctx context.Context, module, action string) bool {
	var (
		ormer                      = s.dbFactory.GetDB()
		maintainConfiguration, err = s.maintainConfigurationRepository.GetByID(ormer, module, action)
	)
	if err != nil && err != orm.ErrNoRows {
		logger.Errorf(ctx, "Getting MaintainConfiguration by ID: %v", err)
		return false
	}
	if maintainConfiguration != nil && maintainConfiguration.Status == models.MaintainStatusDisable {
		return true
	}
	return false
}
