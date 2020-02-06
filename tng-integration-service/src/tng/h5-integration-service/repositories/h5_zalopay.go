package repositories

import (
	"github.com/astaxie/beego/orm"
	"tng/common/models/h5-intergration"
	"tng/common/utils/db"
)

// H5ZalopayRepository represents a repository for managing h5 service of zalopay.
type H5ZalopayRepository interface {
	SetTokenInfo(ormer *db.DB, info *h5_intergration.H5ZaloPay) error
	GetTokenInfo(ormer *db.DB, appID int, userID string) (*h5_intergration.H5ZaloPay, error)
	DeleteTokenInfo(ormer *db.DB, appID int, userID string) error
}

type h5ZaloPayRepository struct{}

// NewH5ZalopayRepository create a new instance of H5ZaloPay Repository.
func NewH5ZalopayRepository(dbFactory db.Factory, ) H5ZalopayRepository {
	return &h5ZaloPayRepository{}
}

func (h *h5ZaloPayRepository) SetTokenInfo(ormer *db.DB, info *h5_intergration.H5ZaloPay) error {
	_, err := ormer.InsertOrUpdate(info)
	return err
}

func (h *h5ZaloPayRepository) GetTokenInfo(ormer *db.DB, appID int, userID string) (*h5_intergration.H5ZaloPay, error) {
	cond := orm.NewCondition().
		And("app_id", appID).
		And("user_id", userID)
	tokenInfo := &h5_intergration.H5ZaloPay{}
	qs := ormer.QueryTable(new(h5_intergration.H5ZaloPay))
	err := qs.SetCond(cond).One(tokenInfo)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func (h *h5ZaloPayRepository) DeleteTokenInfo(ormer *db.DB, appID int, userID string) error {
	_, err := ormer.QueryTable(new(h5_intergration.H5ZaloPay)).
		Filter("app_id", appID).
		Filter("user_id", userID).Delete()
	if err != nil {
		return err
	}
	return nil
}
