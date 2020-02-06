package services

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/dig"
	zaloAdapter "tng/common/adapters/zalo"
	"tng/common/utils/cfgutil"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/user-profile-service/helper"
	"tng/user-profile-service/repositories"
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

	_ = serviceContainer.Provide(func() zaloAdapter.Adapter {
		var (
			baseURL = cfgutil.Load("ZALO_OA_URL")
		)
		return zaloAdapter.NewAdapter(
			baseURL,
		)
	})

	_ = serviceContainer.Provide(db.NewDBFactory)
	_ = serviceContainer.Provide(repositories.NewUserProfileRepository)
	_ = serviceContainer.Provide(repositories.NewUserSessionRepository)
	_ = serviceContainer.Provide(NewUserProfileService)
	_ = serviceContainer.Provide(NewAuthenticationService)
	_ = serviceContainer.Provide(helper.NewHelper)
	return serviceContainer
}

// GetServiceContainer returns a new instance of Service Container
func GetServiceContainer() *dig.Container {
	return serviceContainer
}
