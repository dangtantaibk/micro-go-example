package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	kafka "github.com/Shopify/sarama"
	"time"
)

//Writer ...
type Writer interface {
	WriteRaw([]byte)
	Write(kafka.Encoder)
	WriteRawByTopic([]byte, string)
	Close()
	WriteByTopic(interface{}, string) error
}
type writer struct {
	topics   []string
	producer kafka.AsyncProducer
}

func SyncWrite(cfg KafkaCfg, topic string, v []byte) error {
	cfgKafka := kafka.NewConfig()
	cfgKafka.Producer.RequiredAcks = kafka.WaitForLocal
	cfgKafka.Producer.Return.Successes = true
	if len(KafkaAddress) > 0 {
		cfg.Addrs = KafkaAddress
	}
	producer, err := kafka.NewSyncProducer(cfg.Addrs, cfgKafka)
	if err != nil {
		return err
	}
	_, _, err = producer.SendMessage(&kafka.ProducerMessage{
		Topic: GetTopic(topic),
		Value: sarama.ByteEncoder(v),
	})
	return err

}

func SyncWriteAndClose(cfg KafkaCfg, topic string, v []byte) error {
	if len(KafkaAddress) > 0 {
		cfg.Addrs = KafkaAddress
	}
	cfgKafka := kafka.NewConfig()
	cfgKafka.Producer.RequiredAcks = kafka.WaitForLocal
	cfgKafka.Producer.Return.Successes = true
	cfgKafka.Producer.MaxMessageBytes = 64 * 1000000

	producer, err := kafka.NewSyncProducer(cfg.Addrs, cfgKafka)
	if err != nil {
		return err
	}
	defer producer.Close()
	_, _, err = producer.SendMessage(&kafka.ProducerMessage{
		Topic: GetTopic(topic),
		Value: sarama.ByteEncoder(v),
	})
	return err

}

func CreateWriters(cfg KafkaCfg) Writer {
	return createWriter(cfg.Addrs, cfg.Topics, cfg.MaxMessageBytes, cfg.Compress)
}

func createWriter(addrs, topics []string, maxMessageBytes int, compress bool) Writer {
	if len(KafkaAddress) > 0 {
		addrs = KafkaAddress
	}
	cfg := kafka.NewConfig()
	cfg.Producer.RequiredAcks = kafka.WaitForLocal
	cfg.Producer.Flush.Frequency = 50 * time.Millisecond
	if maxMessageBytes > 0 {
		cfg.Producer.MaxMessageBytes = maxMessageBytes
	}
	if compress {
		cfg.Producer.Compression = kafka.CompressionGZIP
	}

	producer, err := kafka.NewAsyncProducer(addrs, cfg)
	if err != nil {
		return nil
	}
	go func() {
		for _ = range producer.Errors() {
		}
	}()
	return &writer{topics: topics,
		producer: producer}

}
func (w *writer) Write(v kafka.Encoder) {
	if w == nil {
		return
	}
	for _, topic := range w.topics {
		w.producer.Input() <- &kafka.ProducerMessage{
			Topic: GetTopic(topic),
			Value: v,
		}
	}
}
func (w *writer) Close() {
	w.producer.AsyncClose()
}
func (w *writer) WriteRaw(v []byte) {
	if w == nil {
		return
	}
	for _, topic := range w.topics {

		w.producer.Input() <- &kafka.ProducerMessage{
			Topic: GetTopic(topic),
			Value: sarama.ByteEncoder(v),
		}
	}
}

func (w *writer) WriteRawByTopic(v []byte, topicName string) {
	if w == nil {
		return
	}
	w.producer.Input() <- &kafka.ProducerMessage{
		Topic: GetTopic(topicName),
		Value: sarama.ByteEncoder(v),
	}
}
func (w *writer) WriteByTopic(v interface{}, topicName string) error {
	if data, err := json.Marshal(v); err == nil {
		w.producer.Input() <- &kafka.ProducerMessage{
			Topic: GetTopic(topicName),
			Value: sarama.ByteEncoder(data),
		}
		return nil
	} else {
		return err
	}
}
