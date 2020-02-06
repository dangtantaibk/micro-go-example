package mqttcli

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"net/url"
	"time"
)

type callBack func(topic, payload string)

type MqttCli interface {
	Pub(ctx context.Context, topic string, msg string)
	Sub(ctx context.Context, topic string, cb callBack)
}

type mqttCli struct {
	client mqtt.Client
	uri    url.URL
}

func NewMqtt(uri url.URL) (MqttCli, error) {
	clientId := uuid.New().String()
	client, err := connect(clientId, uri)
	if err != nil {
		return nil, err
	}
	return &mqttCli{
		client: client,
	}, nil
}

func connect(clientId string, uri url.URL) (mqtt.Client, error) {
	opts := clientOptions(clientId, &uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		return nil, err
	}
	return client, nil
}

func (m *mqttCli) Pub(ctx context.Context, topic string, msg string) {
	m.client.Publish(topic, 0, false, msg)
}

func (m *mqttCli) Sub(ctx context.Context, topic string, cb callBack) {
	m.client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		cb(topic, string(msg.Payload()))
	})
}

func clientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri.Host)
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	if password != "" {
		opts.SetPassword(password)
	}
	opts.SetClientID(clientId)
	return opts
}
