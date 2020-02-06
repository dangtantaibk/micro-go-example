package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
	"tng/common/utils/cfgutil"

	"tng/common/concurrency"
)

var (
	servicePrefixName    = cfgutil.Load("APPNAME")
	KafkaAddress         = getAddress()
	once                 sync.Once
	kafkaErrorList       = concurrency.NewConcurrentSlice(20)
	kafkaErrorPrefixName = servicePrefixName
	MsgSubjectKafkaError = "Thông báo lỗi hệ thống kafka"
)

type KafkaCfg struct {
	Addrs           []string
	Topics          []string
	Group           string
	Oldest          bool
	MaxMessageBytes int
	Compress        bool
	Newest          bool
}

func GetTopic(topic string) string {
	if servicePrefixName != "" && strings.HasPrefix(topic, servicePrefixName) == false {
		return servicePrefixName + topic
	}
	return topic
}
func getAddress() []string {
	addr := os.Getenv("KAFKA_ADDR")
	if len(addr) > 0 {
		return strings.Split(addr, ",")
	}
	return nil
}

func SetKafkaErrorConfig(prefixName string) {
	kafkaErrorPrefixName = strings.Replace(prefixName, ".", "", -1)
}

func addKafkaError(address []string, err error) {
	kafkaErrorList.Append(strings.Join(address, ",") + ": " + err.Error())
}

func processKafkaError() {
	once.Do(func() {
		go func() {
			signals := make(chan os.Signal, 1)
			signal.Notify(signals, os.Interrupt)
			tick := time.Tick(5 * time.Minute)
			for {
				select {
				case <-tick:
					errText := kafkaErrorList.ToStringAndClear()
					if len(errText) > 0 {
						sendKafkaAlert(errText)
					}
				case <-signals:
					return
				}
			}
		}()
	})
}

func sendKafkaAlert(errText string) {
	if len(kafkaErrorPrefixName) > 0 {
		MsgSubjectKafkaError = fmt.Sprintf("[%s] %s", kafkaErrorPrefixName, MsgSubjectKafkaError)
	}
}
