package repositories

import (
	"github.com/astaxie/beego/orm"
	"tng/common/models/shiper"
	"tng/common/utils/db"
)

type UserRepository interface {
	Insert(ormer *db.DB, user *shiper.User) error
	GetByPhoneNumber(ormer *db.DB, phoneNumber string) (*shiper.Shipper, error)
	GetByPhoneNumberAndPassword(ormer *db.DB, phoneNumber string, password string) (*shiper.User, error)
}

type userRepository struct{}

func (u *userRepository) GetByPhoneNumberAndPassword(ormer *db.DB, phoneNumber string, password string) (*shiper.User, error) {
	cond := orm.NewCondition().
		And("phone_number", phoneNumber).
		And("password", password)
	userInfo := &shiper.User{}
	qs := ormer.QueryTable(new(shiper.User))
	err := qs.SetCond(cond).One(userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (u *userRepository) GetByPhoneNumber(ormer *db.DB, phoneNumber string) (*shiper.Shipper, error) {
	cond := orm.NewCondition().
		And("phone", phoneNumber)
	shipperInfo := &shiper.Shipper{}
	qs := ormer.QueryTable(new(shiper.Shipper))
	err := qs.SetCond(cond).One(shipperInfo)
	if err != nil {
		return nil, err
	}
	return shipperInfo, nil
}


func (u *userRepository) Insert(ormer *db.DB, user *shiper.User) error {
	_, err := ormer.Insert(user)
	return err
}

func NewUserRepository(dbFactory db.Factory) UserRepository {
	return &userRepository{}
}