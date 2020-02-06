package kafka

import (
	"os"
	"os/signal"
	"time"

	"tng/common/concurrency"

	"github.com/Shopify/sarama"
	kafka "github.com/bsm/sarama-cluster"
	"github.com/go-redis/redis"
)

type ReadCB func([]byte)
type ReadExCB func(string, []byte)
type ReadCBWithErr func(string, []byte) error

type Reader interface {
	Read(cb ReadCB)
	ReadEx(cb ReadExCB)
	ReadWithErr(cb ReadCBWithErr, q *redis.Client)
}
type reader struct {
	consumer *kafka.Consumer
	topics   []string
	group    string
}

func CreateReaders(cfg KafkaCfg) Reader {
	processKafkaError()
	config := kafka.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	if cfg.Newest {
		config.Consumer.Offsets.Initial = sarama.OffsetNewest
	} else {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}
	if len(KafkaAddress) > 0 {
		cfg.Addrs = KafkaAddress
	}
	topics := []string{}
	for _, topic := range cfg.Topics {
		topics = append(topics, GetTopic(topic))
	}
	c, err := kafka.NewConsumer(cfg.Addrs, cfg.Group, topics, config)

	if err != nil {
		return &reader{}
	}
	go func() {
		for err := range c.Errors() {
			addKafkaError(cfg.Addrs, err)
		}
	}()

	go func() {
		for _ = range c.Notifications() {
		}
	}()
	return &reader{consumer: c, topics: topics, group: cfg.Group}
}

func (r *reader) Read(cb ReadCB) {
	if r.consumer == nil {
		return
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	consumer := r.consumer
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				cb(msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed

			}
		case <-signals:
			return
		}
	}

}

func (r *reader) ReadEx(cb ReadExCB) {
	if r.consumer == nil {
		return
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	consumer := r.consumer
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				cb(msg.Topic, msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case <-signals:
			return
		}
	}
}
func (r *reader) retryMessage(cb ReadCBWithErr, q *redis.Client) {
	cc := concurrency.New()
	for _, topic := range r.topics {
		_topic := topic
		key := r.group + topic
		cc.Add(func() error {
			datas, err := q.LRange(key, 0, -1).Result()
			if err != nil {
				return nil
			}
			if len(datas) <= 0 {
				return nil
			}
			q.Del(key)
			for _, data := range datas {
				if err = cb(_topic, []byte(data)); err != nil {
				}
			}
			return nil
		})
	}
	cc.Do()

}
func (r *reader) ReadWithErr(cb ReadCBWithErr, q *redis.Client) {
	if r.consumer == nil {
		return
	}
	done := make(chan bool, 1)
	if q != nil {
		r.retryMessage(cb, q)
		go func() {
			for {
				select {
				case <-done:
					return
				case <-time.After(time.Hour):
					{
						r.retryMessage(cb, q)
					}
				}

			}
		}()
	}

	const nWoker = 64
	consumer := r.consumer
	messages := consumer.Messages()
	for i := 0; i < nWoker; i++ {
		go func() {
			for msg := range messages {
				if err := cb(msg.Topic, msg.Value); err != nil {
					key := r.group + msg.Topic
					if q != nil {
						if n, err := q.LLen(key).Result(); err == nil && n < 4096 {
							q.LPush(key, msg.Value)
						} else {
						}
					}
				}
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		}()
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
	done <- true
	consumer.Close()
}
