package nsq

import "github.com/nsqio/go-nsq"

func NewNsqConsumer(channel, secretKey, topic string) (*nsq.Consumer, error) {
	nsqCFG := nsq.NewConfig()
	nsqCFG.AuthSecret = secretKey

	return nsq.NewConsumer(topic, channel, nsqCFG)
}
