package nsq

import "github.com/nsqio/go-nsq"

func NewNsqPublisher(host, port, secretKey string) (*nsq.Producer, error) {
	nsqCFG := nsq.NewConfig()
	nsqCFG.AuthSecret = secretKey

	return nsq.NewProducer(host+":"+port, nsqCFG)
}
