package repositories

import (
	"github.com/astaxie/beego/orm"
	"tng/common/models/shiper"
	"tng/common/utils/db"
)

type shipperRepository struct{}

func (s *shipperRepository) Delete(ormer *db.DB, shipperId int64) error {
	_, err := ormer.Delete(&shiper.Shipper{ID:shipperId})
	if err != nil {
		return err
	}
	return nil
}

func (s *shipperRepository) Update(ormer *db.DB, shipper *shiper.Shipper) (*shiper.Shipper, error) {
	_, err := ormer.Update(shipper)
	if err != nil {
		return nil, err
	}
	qs := ormer.QueryTable(new(shiper.Shipper))
	cond := orm.NewCondition().And("id", shipper.ID)
	shipperInfo := &shiper.Shipper{}
	err = qs.SetCond(cond).One(shipperInfo)
	if err != nil {
		return nil, err
	}
	return shipperInfo, nil
}

func (s *shipperRepository) List(ormer *db.DB, current int32, totalPerPage int32) ([] *shiper.Shipper, int64, error) {
	var (
		list []*shiper.Shipper
		qs1 = ormer.QueryTable(new(shiper.Shipper))
		qs2 = ormer.QueryTable(new(shiper.Shipper))
		fromRecord int32
		totalRecord int64
	)
	fromRecord = current * totalPerPage - totalPerPage

	totalRecord, err := qs1.Count()
	if err != nil {
		return nil, 0, err
	}
	qs2 = qs2.Limit(totalPerPage, fromRecord)
	if _, err := qs2.All(&list); err != nil {
		return nil, 0, err
	}
	return list, totalRecord, nil
}

func (s *shipperRepository) Insert(ormer *db.DB, shipper *shiper.Shipper) error {
	_, err := ormer.Insert(shipper)
	return err
}

func (s *shipperRepository) GetByPhoneNumber(ormer *db.DB, phoneNumber string) (*shiper.Shipper, error) {
	cond := orm.NewCondition().
		And("phone_number", phoneNumber)
	shipperInfo := &shiper.Shipper{}
	qs := ormer.QueryTable(new(shiper.Shipper))
	err := qs.SetCond(cond).One(shipperInfo)
	if err != nil {
		return nil, err
	}
	return shipperInfo, nil
}

func (s *shipperRepository) GetByPhoneNumberAndPassword(ormer *db.DB, phoneNumber string, password string) (*shiper.Shipper, error) {
	cond := orm.NewCondition().
		And("phone_number", phoneNumber).
		And("password", password)
	shipperInfo := &shiper.Shipper{}
	qs := ormer.QueryTable(new(shiper.Shipper))
	err := qs.SetCond(cond).One(shipperInfo)
	if err != nil {
		return nil, err
	}
	return shipperInfo, nil
}

type ShipperRepository interface {
	Insert(ormer *db.DB, shipper *shiper.Shipper) error
	GetByPhoneNumber(ormer *db.DB, phoneNumber string) (*shiper.Shipper, error)
	GetByPhoneNumberAndPassword(ormer *db.DB, phoneNumber string, password string) (*shiper.Shipper, error)
	List(ormer *db.DB, current int32, totalPerPage int32) ([] *shiper.Shipper, int64, error)
	Update(ormer *db.DB, shipper *shiper.Shipper) (*shiper.Shipper, error)
	Delete(ormer *db.DB, shipperId int64) error
}

func NewShipperRepository(dbFactory db.Factory) ShipperRepository {
	return &shipperRepository{}
}
