package kafka

import (
	"bitbucket.org/ssinbeti/crm-common/pkg/config"
	"github.com/segmentio/kafka-go"
	"time"
)

func GetReader(cfg config.KafkaConfigs) *kafka.Reader {
	rConfig := kafka.ReaderConfig{
		Brokers:   cfg.Brokers,
		GroupID:   cfg.GroupID,
		Topic:     cfg.Topic,
		Partition: 0,
		Dialer: &kafka.Dialer{
			Timeout:   10 * time.Second,
			DualStack: true,
		},
	}

	return kafka.NewReader(rConfig)
}
