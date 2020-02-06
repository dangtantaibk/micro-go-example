package services

import (
	"go.uber.org/dig"
	"net/url"
	_h5ZaloPayAdapter "tng/common/adapters/h5-zalopay"
	"tng/common/utils/cfgutil"
	"tng/common/utils/db"
	"tng/common/utils/mqttcli"
	"tng/dev-tool-service/repositories"
)

// serviceContainer is a global ServiceProvider.
var serviceContainer *dig.Container

// InitialServices inits service provider.
func InitialServices() *dig.Container {
	serviceContainer = dig.New()

	_ = serviceContainer.Provide(func() _h5ZaloPayAdapter.Adapter {
		var (
			baseURL     = cfgutil.Load("H5_ZALOPAY_URL")
			clientID, _ = cfgutil.LoadInt("H5_ZALOPAY_CLIENT_ID")
			clientKey   = cfgutil.Load("H5_ZALOPAY_CLIENT_KEY")
		)
		return _h5ZaloPayAdapter.NewAdapter(
			baseURL,
			clientKey,
			int64(clientID),
		)
	})

	_ = serviceContainer.Provide(func() mqttcli.MqttCli {
		var (
			mqttHost = cfgutil.Load("MQTT_HOST")
			mqttPath = cfgutil.Load("MQTT_PATH")
			uri      url.URL
		)
		uri.Host = mqttHost
		uri.Path = mqttPath
		clientMqtt, err := mqttcli.NewMqtt(uri)

		if err != nil {
			panic(err)
		}
		return clientMqtt
	})

	_ = serviceContainer.Provide(db.NewDBFactory)
	_ = serviceContainer.Provide(repositories.NewMaintainConfigurationRepository)
	_ = serviceContainer.Provide(repositories.NewPi3UpdateInfoRepository)
	_ = serviceContainer.Provide(NewPi3UpdateInfoService)
	return serviceContainer
}

// GetServiceContainer returns a new instance of Service Container
func GetServiceContainer() *dig.Container {
	return serviceContainer
}
