package grpc

import (
	"bitbucket.org/ssinbeti/crm-common/pkg/config"
	"github.com/cenkalti/backoff"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

func NewGrpcClient(config config.GRPCClient) (*grpc.ClientConn, error) {
	initializeTimeout, err := time.ParseDuration(config.ConnectionBackoffTimeout)
	if err != nil {
		return nil, err
	}

	return newConnection(initializeTimeout, config.Host, config.Port)
}

func newConnection(timeout time.Duration, addr string, port int) (conn *grpc.ClientConn, err error) {
	address := addr + ":" + strconv.Itoa(port)
	operation := func() error {
		conn, err = grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Printf("did not connect: %s", err)
		}
		return nil
	}
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = timeout
	err = backoff.Retry(operation, b)

	log.Println("Successfully connected to grpc server, address: ", address)
	return conn, err
}
