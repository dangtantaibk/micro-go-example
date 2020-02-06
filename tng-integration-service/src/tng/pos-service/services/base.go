package services

import (
	"github.com/patrickmn/go-cache"

	"tng/common/utils/db"
)

// BaseService keep the basic feature for all services.
type BaseService struct {
	dbFactory  db.Factory
	localCache *cache.Cache
}
