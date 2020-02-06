package repositories

import (
	"github.com/astaxie/beego/orm"
	"tng/common/models/menu"
	"tng/common/utils/db"
)

type CategoryRepository interface {
	Insert(ormer *db.DB, category *menu.Category) error
	List(ormer *db.DB) ([] *menu.Category, error)
	UpdateStatus(ormer *db.DB, categoryId int64, status int) error
	Delete(ormer *db.DB, categoryId int64) error
	Update(ormer *db.DB, category *menu.Category) error
}

type categoryRepository struct {
	
}

func (c *categoryRepository) Update(ormer *db.DB, category *menu.Category) error {
	_, err := ormer.Update(category)
	return err
}

func (c *categoryRepository) Delete(ormer *db.DB, categoryId int64) error {
	_, err := ormer.Delete(&menu.Category{ID:categoryId})
	if err != nil {
		return err
	}
	return nil
}

func (c *categoryRepository) UpdateStatus(ormer *db.DB, categoryId int64, status int) error {
	qs := ormer.QueryTable(new(menu.Category))
	cond := orm.NewCondition().And("category_id", categoryId)
	var categoryInfo menu.Category
	err := qs.SetCond(cond).One(&categoryInfo)
	if err != nil {
		return err
	}
	categoryInfo.Status = status
	_, err = ormer.Update(&categoryInfo)
	if err != nil {
		return err
	}
	return nil
}

func (c *categoryRepository) List(ormer *db.DB) ([] *menu.Category, error) {
	var(
		list []*menu.Category
		qs = ormer.QueryTable(new(menu.Category))
	)

	if _, err := qs.All(&list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *categoryRepository) Insert(ormer *db.DB, category *menu.Category) error {
	_, err := ormer.Insert(category)
	return err
}

func NewCategoryRepository(factory db.Factory) CategoryRepository {
	return &categoryRepository{}
}