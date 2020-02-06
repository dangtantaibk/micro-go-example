package repositories

import (
	"tng/common/models/menu"
	"tng/common/utils/db"
)

type areaRepository struct {
	
}

func (a areaRepository) List(ormer *db.DB) ([] *menu.Area, error) {
	var (
		list []*menu.Area
		qs = ormer.QueryTable(new(menu.Area))
	)

	if _, err := qs.All(&list); err != nil {
		return nil, err
	}
	return list, nil
}

type AreaRepository interface {
	List(ormer *db.DB) ([] *menu.Area, error)
}

func NewAreaRepository() AreaRepository {
	return &areaRepository{}
}