package repositories

import (
	"github.com/astaxie/beego/orm"
	userProfileModel "tng/common/models/user-profile"
	"tng/common/utils/db"
)

type UserProfileRepository interface {
	Create(*db.DB, *userProfileModel.User) (int64, error)
	Update(*db.DB, *userProfileModel.User) error
	GetByID(*db.DB, string) (*userProfileModel.User, error)
	GetByUCode(*db.DB, string) (*userProfileModel.User, error)
	GetBySocialID(*db.DB, string) (*userProfileModel.User, error)
}

type userProfileRepository struct{}

func NewUserProfileRepository() UserProfileRepository {
	return &userProfileRepository{}
}

func (r *userProfileRepository) Create(ormer *db.DB, dataInput *userProfileModel.User) (int64, error) {
	return ormer.Insert(dataInput)
}

func (r *userProfileRepository) Update(ormer *db.DB, dataInput *userProfileModel.User) error {
	_, err := ormer.Update(dataInput)
	return err
}

func (r *userProfileRepository) GetByID(ormer *db.DB, userID string) (*userProfileModel.User, error) {
	cond := orm.NewCondition().
		And("id", userID)
	userInfo := &userProfileModel.User{}
	qs := ormer.QueryTable(new(userProfileModel.User))
	err := qs.SetCond(cond).One(userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (r *userProfileRepository) GetByUCode(ormer *db.DB, uCode string) (*userProfileModel.User, error) {
	cond := orm.NewCondition().
		And("ucode", uCode)
	userInfo := &userProfileModel.User{}
	qs := ormer.QueryTable(new(userProfileModel.User))
	err := qs.SetCond(cond).One(userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (r *userProfileRepository) GetBySocialID(ormer *db.DB, socialID string) (*userProfileModel.User, error) {
	cond := orm.NewCondition().
		And("social_id", socialID)
	userInfo := &userProfileModel.User{}
	qs := ormer.QueryTable(new(userProfileModel.User))
	err := qs.SetCond(cond).One(userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
