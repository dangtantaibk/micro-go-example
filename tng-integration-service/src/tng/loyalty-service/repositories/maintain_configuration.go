package repositories

import (
	"tng/common/models"
	"tng/common/utils/db"
)

// MaintainConfigurationRepository provides a access to MaintainConfiguration store.
type MaintainConfigurationRepository interface {
	GetByID(ormer *db.DB, module, action string) (*models.MaintainConfiguration, error)
}

type maintainConfigurationRepository struct{}

// NewMaintainConfigurationRepository returns a new instance of MaintainConfigurationRepository.
func NewMaintainConfigurationRepository() MaintainConfigurationRepository {
	return &maintainConfigurationRepository{}
}

func (r *maintainConfigurationRepository) GetByID(ormer *db.DB, module, action string) (*models.MaintainConfiguration, error) {
	var maintainConfiguration models.MaintainConfiguration
	qs := ormer.QueryTable(new(models.MaintainConfiguration)).
		Filter("module", module).
		Filter("action", action).
		Filter("deleted_at__isnull", true)
	if err := qs.One(&maintainConfiguration); err != nil {
		return nil, err
	}
	return &maintainConfiguration, nil
}
