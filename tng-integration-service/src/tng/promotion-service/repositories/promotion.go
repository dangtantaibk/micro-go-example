package repositories

import (
	"tng/common/models/promotion"
	"tng/common/utils/db"
)

// PromotionRepository represents a repository.
type PromotionRepository interface {
	Insert(ormer *db.DB, info *promotion.Campaign) error
}

type promotionRepository struct{}

// NewPromotionRepository create a new instance of Repository.
func NewPromotionRepository(dbFactory db.Factory, ) PromotionRepository {
	return &promotionRepository{}
}

func (h *promotionRepository) Insert(ormer *db.DB, info *promotion.Campaign) error {
	_, err := ormer.InsertOrUpdate(info)
	return err
}
