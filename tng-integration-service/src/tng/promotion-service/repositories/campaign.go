package repositories

import (
	"github.com/astaxie/beego/orm"
	"tng/common/models/promotion"
	"tng/common/utils/db"
	"tng/promotion-service/dtos"
)

// CampaignRepository represents a repository.
type CampaignRepository interface {
	Insert(*db.DB, *promotion.Campaign) error
	List(*db.DB, *dtos.ListCampaignRequest) ([]*promotion.Campaign, error)
}

type campaignRepository struct{}

// NewCampaignRepository create a new instance of Repository.
func NewCampaignRepository(dbFactory db.Factory) CampaignRepository {
	return &campaignRepository{}
}

func (h *campaignRepository) Insert(ormer *db.DB, info *promotion.Campaign) error {
	_, err := ormer.InsertOrUpdate(info)
	return err
}

func (h *campaignRepository) List(ormer *db.DB, request *dtos.ListCampaignRequest) ([]*promotion.Campaign, error) {
	var (
		list      []*promotion.Campaign
		qs        = ormer.QueryTable(new(promotion.Campaign))
		cond      = orm.NewCondition()
	)
	cond = cond.And("appid", request.AppID)
	qs = qs.SetCond(cond)
	qs = qs.Limit(request.PageSize).Offset((request.PageIndex - 1) * request.PageSize)
	if _, err := qs.All(&list); err != nil {
		return nil, err
	}
	return list, nil
}