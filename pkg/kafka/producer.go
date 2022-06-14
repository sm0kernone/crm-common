package kafka

import (
	"github.com/sm0kernone/crm-common/pkg/config"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	"time"
)

func GetKafkaConn(cfg config.KafkaConfigs) *kafka.Writer {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: cfg.GroupID,
	}

	config := kafka.WriterConfig{
		Brokers:          []string{cfg.BrokerProducer},
		Topic:            cfg.Topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}

	//conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.BrokerProducer, cfg.Topic, 0)

	//if err != nil {
	//	return nil, err
	//}

	//conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	return kafka.NewWriter(config)
}
