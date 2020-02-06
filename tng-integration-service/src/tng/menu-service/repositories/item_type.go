package repositories

import (
	"github.com/astaxie/beego/orm"
	"tng/common/models/menu"
	"tng/common/utils/db"
)

type itemTypeRepository struct {

}

func (a itemTypeRepository) Update(ormer *db.DB, itemType *menu.ItemType) (*menu.ItemType, error) {
	_, err := ormer.Update(itemType)
	if err != nil {
		return nil, err
	}
	qs := ormer.QueryTable(new(menu.ItemType))
	cond := orm.NewCondition().And("item_type_id", itemType.ID)
	itemTypeInfo := &menu.ItemType{}
	err = qs.SetCond(cond).One(itemTypeInfo)
	if err != nil {
		return nil, err
	}
	return itemTypeInfo, nil
}

func (a itemTypeRepository) Insert(ormer *db.DB, itemType *menu.ItemType) (*menu.ItemType, error) {
	id, err := ormer.Insert(itemType)
	if err != nil {
		return nil, err
	}
	qs := ormer.QueryTable(new(menu.ItemType))
	cond := orm.NewCondition().And("item_type_id", id)
	//var itemTypeInfo *menu.ItemType
	itemTypeInfo := &menu.ItemType{}
	err = qs.SetCond(cond).One(itemTypeInfo)
	if err != nil {
		return nil, err
	}
	return itemTypeInfo, nil
}

func (a itemTypeRepository) Delete(ormer *db.DB, itemTypeId int64) error {
	_, err := ormer.Delete(&menu.ItemType{ID:itemTypeId})
	if err != nil {
		return err
	}
	return nil
}

func (a itemTypeRepository) List(ormer *db.DB) ([] *menu.ItemType, error) {
	var (
		list []*menu.ItemType
		qs = ormer.QueryTable(new(menu.ItemType))
	)

	if _, err := qs.All(&list); err != nil {
		return nil, err
	}
	return list, nil
}

type ItemTypeRepository interface {
	List(ormer *db.DB) ([] *menu.ItemType, error)
	Delete(ormer *db.DB, categoryId int64) error
	Insert(ormer *db.DB, itemType *menu.ItemType) (*menu.ItemType, error)
	Update(ormer *db.DB, itemType *menu.ItemType) (*menu.ItemType, error)
}

func NewItemTypeRepository() ItemTypeRepository {
	return &itemTypeRepository{}
}