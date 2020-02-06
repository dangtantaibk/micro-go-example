package services

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/dig"
	"tng/common/utils/cfgutil"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/promotion-service/repositories"
)

// serviceContainer is a global ServiceProvider.
var serviceContainer *dig.Container

// InitialServices inits service provider.
func InitialServices() *dig.Container {
	serviceContainer = dig.New()

	_ = serviceContainer.Provide(func() redisutil.Cache {
		var (
			redisDB, _ = cfgutil.LoadInt("REDIS_DB")
			redisHost  = cfgutil.Load("REDIS_HOST")
			redisPort  = cfgutil.Load("REDIS_PORT")
		)
		redisClient := redis.NewClient(
			&redis.Options{
				Addr:     fmt.Sprintf("%v:%v", redisHost, redisPort),
				Password: cfgutil.Load("REDIS_PASSWORD"),
				DB:       redisDB,
			})
		pong := redisClient.Ping()
		if _, err := pong.Result(); err != nil {
			panic(err)
		}
		redisCli, err := redisutil.NewCache(redisClient)
		if err != nil {
			panic(err)
		}
		return redisCli
	})

	_ = serviceContainer.Provide(db.NewDBFactory)
	_ = serviceContainer.Provide(repositories.NewPromotionRepository)
	_ = serviceContainer.Provide(repositories.NewCampaignRepository)
	_ = serviceContainer.Provide(NewPromotionService)
	_ = serviceContainer.Provide(NewCampaignService)
	return serviceContainer
}

// GetServiceContainer returns a new instance of Service Container
func GetServiceContainer() *dig.Container {
	return serviceContainer
}
