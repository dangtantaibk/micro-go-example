
package repositories

import (
	"context"
	"github.com/astaxie/beego/orm"
	"tng/common/models/loyalty"
	"tng/common/utils/db"
)

type WalletRepository interface {
	InsertOrUpdate(context.Context, *db.DB, *loyalty.Wallet) error
	Delete(context.Context, *db.DB, int64) error
	GetByID(context.Context, *db.DB, int64) (*loyalty.Wallet, error)
	List(context.Context, *db.DB, int32, int32) ([]*loyalty.Wallet, error)
	GetByUserID(context.Context, *db.DB, string) (*loyalty.Wallet, error)
}

type walletRepository struct{}

func NewWalletRepository(dbFactory db.Factory) WalletRepository {
	return &walletRepository{}
}

func (r *walletRepository) InsertOrUpdate(ctx context.Context, ormer *db.DB, dataInput *loyalty.Wallet) error {
	_, err := ormer.InsertOrUpdate(dataInput)
	return err
}

func (r *walletRepository) Delete(ctx context.Context, ormer *db.DB, id int64) error {
	qs := ormer.QueryTable(new(loyalty.Wallet))
	cond := orm.NewCondition().
		And("id", id)
	_, err := qs.SetCond(cond).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r *walletRepository) GetByID(ctx context.Context, ormer *db.DB, id int64) (*loyalty.Wallet, error) {
	cond := orm.NewCondition().
		And("id", id)
	item := &loyalty.Wallet{}
	qs := ormer.QueryTable(new(loyalty.Wallet))
	err := qs.SetCond(cond).One(item)
	if err != nil {
		return nil, err
	}
	return item, nil

}

func (r *walletRepository) List(ctx context.Context, ormer *db.DB, pageIndex int32, pageSize int32) ([]*loyalty.Wallet, error) {
	var (
		list       []*loyalty.Wallet                
		qs         = ormer.QueryTable(new(loyalty.Wallet))
		fromRecord int32
	)
	fromRecord = pageIndex*pageSize - pageSize
	qs = qs.Limit(pageSize, fromRecord)
	if _, err := qs.All(&list); err != nil {
		return nil, err
	}

	return list, nil
}
func (r *walletRepository) GetByUserID(ctx context.Context, ormer *db.DB, user_id string) (*loyalty.Wallet, error) {
	cond := orm.NewCondition().
		And("user_id", user_id)
	item := &loyalty.Wallet{}
	qs := ormer.QueryTable(new(loyalty.Wallet))
	c, err := qs.SetCond(cond).Count()
	if err == nil && c > 0{
		err = qs.SetCond(cond).One(item)
		if err == nil {
			return item, nil
		}
		return nil, err
	}
	return item, nil
}
