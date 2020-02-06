package services

import (
	"go.uber.org/dig"
	"net/url"
	"tng/common/utils/cfgutil"
	"tng/common/utils/db"
	"tng/common/utils/mqttcli"
	"tng/cron-job-service/repositories"
)

// serviceContainer is a global ServiceProvider.
var serviceContainer *dig.Container

// InitialServices inits service provider.
func InitialServices() *dig.Container {
	serviceContainer = dig.New()

	_ = serviceContainer.Provide(func() mqttcli.MqttCli {
		var (
			mqttHost = cfgutil.Load("MQTT_HOST")
			mqttPath = cfgutil.Load("MQTT_PATH")
			mqttUser = cfgutil.Load("MQTT_USER")
			mqttPass = cfgutil.Load("MQTT_PASS")
			uri      url.URL
		)
		uri.Host = mqttHost
		uri.Path = mqttPath
		userInfo := url.UserPassword(mqttUser, mqttPass)
		uri.User = userInfo

		clientMqtt, err := mqttcli.NewMqtt(uri)

		if err != nil {
			panic(err)
		}
		return clientMqtt
	})

	_ = serviceContainer.Provide(db.NewDBFactory)
	_ = serviceContainer.Provide(repositories.NewScheduleRepository)
	_ = serviceContainer.Provide(repositories.NewItemRepository)
	_ = serviceContainer.Provide(NewMenuScheduleService)
	return serviceContainer
}

// GetServiceContainer returns a new instance of Service Container
func GetServiceContainer() *dig.Container {
	return serviceContainer
}
